package main

import (
	"fmt"
	"os"

	"github.com/gnemade360/go-cli"
)

func main() {
	rootCmd := gocli.NewCommand(
		gocli.WithName("greet"),
		gocli.WithShort("A simple greeting application"),
		gocli.WithLong("Greet demonstrates basic go-cli usage with argument validation."),
		gocli.WithArgValidator(gocli.ExactArgs(1)),
		gocli.WithRun(func(cmd *gocli.Command, args []string) error {
			fmt.Printf("Hello, %s! Welcome to go-cli.\n", args[0])
			return nil
		}),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
