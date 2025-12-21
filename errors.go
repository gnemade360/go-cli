package gocli

import (
	"fmt"
	"strings"
)

type InvalidArgsError struct {
	Expected string
	Received int
}

func (e *InvalidArgsError) Error() string {
	return fmt.Sprintf("invalid number of arguments: expected %s, received %d", e.Expected, e.Received)
}

type InvalidArgError struct {
	Arg       string
	ValidArgs []string
}

func (e *InvalidArgError) Error() string {
	if len(e.ValidArgs) == 0 {
		return fmt.Sprintf("invalid argument %q", e.Arg)
	}
	return fmt.Sprintf("invalid argument %q, valid arguments are: %s", e.Arg, strings.Join(e.ValidArgs, ", "))
}
