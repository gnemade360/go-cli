package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gnemade360/go-cli"
)

func main() {
	rootCmd := gocli.NewCommand(
		gocli.WithUse("toolbox"),
		gocli.WithShort("A CLI toolbox with multiple commands"),
		gocli.WithLong("Toolbox demonstrates go-cli's subcommand capabilities."),
	)

	versionCmd := gocli.NewCommand(
		gocli.WithUse("version"),
		gocli.WithShort("Print version information"),
		gocli.WithRun(func(cmd *gocli.Command, args []string) error {
			fmt.Println("toolbox v1.0.0")
			fmt.Println("Built with go-cli")
			return nil
		}),
	)

	echoCmd := gocli.NewCommand(
		gocli.WithUse("echo"),
		gocli.WithShort("Echo the provided arguments"),
		gocli.WithArgs(gocli.MinimumNArgs(1)),
		gocli.WithRun(func(cmd *gocli.Command, args []string) error {
			fmt.Println(strings.Join(args, " "))
			return nil
		}),
	)

	uppercaseCmd := gocli.NewCommand(
		gocli.WithUse("uppercase"),
		gocli.WithShort("Convert text to uppercase"),
		gocli.WithArgs(gocli.MinimumNArgs(1)),
		gocli.WithRun(func(cmd *gocli.Command, args []string) error {
			text := strings.Join(args, " ")
			fmt.Println(strings.ToUpper(text))
			return nil
		}),
	)

	lowercaseCmd := gocli.NewCommand(
		gocli.WithUse("lowercase"),
		gocli.WithShort("Convert text to lowercase"),
		gocli.WithArgs(gocli.MinimumNArgs(1)),
		gocli.WithRun(func(cmd *gocli.Command, args []string) error {
			text := strings.Join(args, " ")
			fmt.Println(strings.ToLower(text))
			return nil
		}),
	)

	rootCmd.AddCommand(versionCmd, echoCmd, uppercaseCmd, lowercaseCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
