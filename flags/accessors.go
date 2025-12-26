package flags

import (
	"fmt"
	"time"
)

func (fs *FlagSet) GetString(name string) (string, error) {
	def, ok := fs.schema[name]
	if !ok {
		return "", fmt.Errorf("flag %s not defined", name)
	}
	if def.Type != FlagString {
		return "", fmt.Errorf("flag %s is not a string", name)
	}

	val := fs.Get(name)
	if val == nil {
		return "", nil
	}
	return val.(string), nil
}

func (fs *FlagSet) GetInt(name string) (int, error) {
	def, ok := fs.schema[name]
	if !ok {
		return 0, fmt.Errorf("flag %s not defined", name)
	}
	if def.Type != FlagInt {
		return 0, fmt.Errorf("flag %s is not an int", name)
	}

	val := fs.Get(name)
	if val == nil {
		return 0, nil
	}
	return val.(int), nil
}

func (fs *FlagSet) GetInt64(name string) (int64, error) {
	def, ok := fs.schema[name]
	if !ok {
		return 0, fmt.Errorf("flag %s not defined", name)
	}
	if def.Type != FlagInt64 {
		return 0, fmt.Errorf("flag %s is not an int64", name)
	}

	val := fs.Get(name)
	if val == nil {
		return 0, nil
	}
	return val.(int64), nil
}

func (fs *FlagSet) GetFloat64(name string) (float64, error) {
	def, ok := fs.schema[name]
	if !ok {
		return 0, fmt.Errorf("flag %s not defined", name)
	}
	if def.Type != FlagFloat64 {
		return 0, fmt.Errorf("flag %s is not a float64", name)
	}

	val := fs.Get(name)
	if val == nil {
		return 0, nil
	}
	return val.(float64), nil
}

func (fs *FlagSet) GetBool(name string) (bool, error) {
	def, ok := fs.schema[name]
	if !ok {
		return false, fmt.Errorf("flag %s not defined", name)
	}
	if def.Type != FlagBool {
		return false, fmt.Errorf("flag %s is not a bool", name)
	}

	val := fs.Get(name)
	if val == nil {
		return false, nil
	}
	return val.(bool), nil
}

func (fs *FlagSet) GetDuration(name string) (time.Duration, error) {
	def, ok := fs.schema[name]
	if !ok {
		return 0, fmt.Errorf("flag %s not defined", name)
	}
	if def.Type != FlagDuration {
		return 0, fmt.Errorf("flag %s is not a duration", name)
	}

	val := fs.Get(name)
	if val == nil {
		return 0, nil
	}
	return val.(time.Duration), nil
}

func (fs *FlagSet) GetStringSlice(name string) ([]string, error) {
	def, ok := fs.schema[name]
	if !ok {
		return nil, fmt.Errorf("flag %s not defined", name)
	}
	if def.Type != FlagStringSlice {
		return nil, fmt.Errorf("flag %s is not a []string", name)
	}

	val := fs.Get(name)
	if val == nil {
		return nil, nil
	}
	return val.([]string), nil
}

func (fs *FlagSet) GetIntSlice(name string) ([]int, error) {
	def, ok := fs.schema[name]
	if !ok {
		return nil, fmt.Errorf("flag %s not defined", name)
	}
	if def.Type != FlagIntSlice {
		return nil, fmt.Errorf("flag %s is not a []int", name)
	}

	val := fs.Get(name)
	if val == nil {
		return nil, nil
	}
	return val.([]int), nil
}
