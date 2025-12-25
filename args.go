package gocli

import "fmt"

func ExactArgs(n int) ArgsValidator {
	return func(cmd *Command, args []string) error {
		if len(args) != n {
			return &InvalidArgsError{
				Expected: fmt.Sprintf("%d arg(s)", n),
				Received: len(args),
			}
		}
		return nil
	}
}

func MinimumNArgs(n int) ArgsValidator {
	return func(cmd *Command, args []string) error {
		if len(args) < n {
			return &InvalidArgsError{
				Expected: fmt.Sprintf("at least %d arg(s)", n),
				Received: len(args),
			}
		}
		return nil
	}
}

func MaximumNArgs(n int) ArgsValidator {
	return func(cmd *Command, args []string) error {
		if len(args) > n {
			return &InvalidArgsError{
				Expected: fmt.Sprintf("at most %d arg(s)", n),
				Received: len(args),
			}
		}
		return nil
	}
}

func RangeArgs(min, max int) ArgsValidator {
	return func(cmd *Command, args []string) error {
		if len(args) < min || len(args) > max {
			return &InvalidArgsError{
				Expected: fmt.Sprintf("between %d and %d arg(s)", min, max),
				Received: len(args),
			}
		}
		return nil
	}
}

func OnlyValidArgs() ArgsValidator {
	return func(cmd *Command, args []string) error {
		for _, arg := range args {
			if !contains(cmd.allowedArgs, arg) {
				return &InvalidArgError{
					Arg:       arg,
					ValidArgs: cmd.allowedArgs,
				}
			}
		}
		return nil
	}
}

func MatchAll(validators ...ArgsValidator) ArgsValidator {
	return func(cmd *Command, args []string) error {
		for _, validator := range validators {
			if err := validator(cmd, args); err != nil {
				return err
			}
		}
		return nil
	}
}

func contains(slice []string, item string) bool {
	for _, s := range slice {
		if s == item {
			return true
		}
	}
	return false
}
