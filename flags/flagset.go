package flags

import (
	"fmt"
	"strconv"
	"strings"
	"time"
)

type FlagSet struct {
	schema  FlagSchema
	values  map[string]interface{}
	changed map[string]bool
	args    []string
}

func NewFlagSet(schema FlagSchema) *FlagSet {
	return &FlagSet{
		schema:  schema,
		values:  make(map[string]interface{}),
		changed: make(map[string]bool),
	}
}

func (fs *FlagSet) Parse(args []string) error {
	remaining := []string{}
	i := 0

	for i < len(args) {
		arg := args[i]

		if !strings.HasPrefix(arg, "-") {
			remaining = append(remaining, arg)
			i++
			continue
		}

		var flagName string
		var flagValue string
		hasValue := false

		if strings.HasPrefix(arg, "--") {
			flagName = strings.TrimPrefix(arg, "--")
			if idx := strings.Index(flagName, "="); idx != -1 {
				flagValue = flagName[idx+1:]
				flagName = flagName[:idx]
				hasValue = true
			}
		} else {
			flagName = strings.TrimPrefix(arg, "-")
			if len(flagName) > 1 && !strings.Contains(flagName, "=") {
				return fmt.Errorf("invalid short flag: %s", arg)
			}
			if idx := strings.Index(flagName, "="); idx != -1 {
				flagValue = flagName[idx+1:]
				flagName = flagName[:idx]
				hasValue = true
			}
		}

		def, err := fs.findFlag(flagName)
		if err != nil {
			return err
		}

		actualFlagName := fs.getFlagName(flagName)

		if def.Type == FlagBool && !hasValue {
			fs.values[actualFlagName] = true
			fs.changed[actualFlagName] = true
			i++
			continue
		}

		if !hasValue {
			if i+1 >= len(args) {
				return fmt.Errorf("flag --%s requires a value", actualFlagName)
			}
			i++
			flagValue = args[i]
		}

		parsedValue, err := fs.parseValue(def.Type, flagValue)
		if err != nil {
			return fmt.Errorf("invalid value for flag --%s: %w", actualFlagName, err)
		}

		if def.Validate != nil {
			if err := def.Validate(parsedValue); err != nil {
				return fmt.Errorf("validation failed for flag --%s: %w", actualFlagName, err)
			}
		}

		fs.values[actualFlagName] = parsedValue
		fs.changed[actualFlagName] = true
		i++
	}

	if err := fs.checkRequired(); err != nil {
		return err
	}

	fs.args = remaining
	return nil
}

func (fs *FlagSet) findFlag(name string) (*FlagDefinition, error) {
	if def, ok := fs.schema[name]; ok {
		return def, nil
	}

	for _, def := range fs.schema {
		if def.Short == name {
			return def, nil
		}
	}

	return nil, fmt.Errorf("unknown flag: %s", name)
}

func (fs *FlagSet) getFlagName(name string) string {
	if _, ok := fs.schema[name]; ok {
		return name
	}

	for flagName, def := range fs.schema {
		if def.Short == name {
			return flagName
		}
	}

	return name
}

func (fs *FlagSet) parseValue(flagType FlagType, value string) (interface{}, error) {
	switch flagType {
	case FlagString:
		return value, nil

	case FlagInt:
		return strconv.Atoi(value)

	case FlagInt64:
		return strconv.ParseInt(value, 10, 64)

	case FlagFloat64:
		return strconv.ParseFloat(value, 64)

	case FlagBool:
		return strconv.ParseBool(value)

	case FlagDuration:
		return time.ParseDuration(value)

	case FlagStringSlice:
		if value == "" {
			return []string{}, nil
		}
		return strings.Split(value, ","), nil

	case FlagIntSlice:
		if value == "" {
			return []int{}, nil
		}
		parts := strings.Split(value, ",")
		result := make([]int, len(parts))
		for i, part := range parts {
			v, err := strconv.Atoi(strings.TrimSpace(part))
			if err != nil {
				return nil, fmt.Errorf("invalid integer in slice: %s", part)
			}
			result[i] = v
		}
		return result, nil

	default:
		return nil, fmt.Errorf("unsupported flag type: %v", flagType)
	}
}

func (fs *FlagSet) checkRequired() error {
	for name, def := range fs.schema {
		if def.Required && !fs.changed[name] {
			return fmt.Errorf("required flag --%s not provided", name)
		}
	}
	return nil
}

func (fs *FlagSet) Get(name string) interface{} {
	if val, ok := fs.values[name]; ok {
		return val
	}
	if def, ok := fs.schema[name]; ok {
		return def.Default
	}
	return nil
}

func (fs *FlagSet) Changed(name string) bool {
	return fs.changed[name]
}

func (fs *FlagSet) Args() []string {
	return fs.args
}
