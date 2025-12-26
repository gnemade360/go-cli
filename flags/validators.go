package flags

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"
)

type Validator func(value interface{}) error

func Range(min, max int) Validator {
	return func(value interface{}) error {
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("expected int, got %T", value)
		}
		if v < min || v > max {
			return fmt.Errorf("must be between %d and %d, got %d", min, max, v)
		}
		return nil
	}
}

func Positive() Validator {
	return func(value interface{}) error {
		v, ok := value.(int)
		if !ok {
			return fmt.Errorf("expected int, got %T", value)
		}
		if v <= 0 {
			return fmt.Errorf("must be positive, got %d", v)
		}
		return nil
	}
}

func NotEmpty() Validator {
	return func(value interface{}) error {
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("expected string, got %T", value)
		}
		if strings.TrimSpace(v) == "" {
			return errors.New("must not be empty")
		}
		return nil
	}
}

func OneOf(allowed ...string) Validator {
	return func(value interface{}) error {
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("expected string, got %T", value)
		}
		for _, a := range allowed {
			if v == a {
				return nil
			}
		}
		return fmt.Errorf("must be one of: %s", strings.Join(allowed, ", "))
	}
}

func Regex(pattern string) Validator {
	re := regexp.MustCompile(pattern)
	return func(value interface{}) error {
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("expected string, got %T", value)
		}
		if !re.MatchString(v) {
			return fmt.Errorf("must match pattern: %s", pattern)
		}
		return nil
	}
}

func MinLength(min int) Validator {
	return func(value interface{}) error {
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("expected string, got %T", value)
		}
		if len(v) < min {
			return fmt.Errorf("must be at least %d characters, got %d", min, len(v))
		}
		return nil
	}
}

func MaxLength(max int) Validator {
	return func(value interface{}) error {
		v, ok := value.(string)
		if !ok {
			return fmt.Errorf("expected string, got %T", value)
		}
		if len(v) > max {
			return fmt.Errorf("must be at most %d characters, got %d", max, len(v))
		}
		return nil
	}
}

func MinDuration(min time.Duration) Validator {
	return func(value interface{}) error {
		v, ok := value.(time.Duration)
		if !ok {
			return fmt.Errorf("expected time.Duration, got %T", value)
		}
		if v < min {
			return fmt.Errorf("must be at least %v, got %v", min, v)
		}
		return nil
	}
}

func MinItems(min int) Validator {
	return func(value interface{}) error {
		switch v := value.(type) {
		case []string:
			if len(v) < min {
				return fmt.Errorf("must have at least %d items, got %d", min, len(v))
			}
		case []int:
			if len(v) < min {
				return fmt.Errorf("must have at least %d items, got %d", min, len(v))
			}
		default:
			return fmt.Errorf("expected slice, got %T", value)
		}
		return nil
	}
}

func UniqueItems() Validator {
	return func(value interface{}) error {
		v, ok := value.([]string)
		if !ok {
			return fmt.Errorf("expected []string, got %T", value)
		}
		seen := make(map[string]bool)
		for _, item := range v {
			if seen[item] {
				return fmt.Errorf("duplicate item: %s", item)
			}
			seen[item] = true
		}
		return nil
	}
}

func All(validators ...Validator) Validator {
	return func(value interface{}) error {
		for _, validator := range validators {
			if err := validator(value); err != nil {
				return err
			}
		}
		return nil
	}
}
