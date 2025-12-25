package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gnemade360/go-cli"
)

func main() {
	rootCmd := gocli.NewCommand(
		gocli.WithName("toolbox"),
		gocli.WithShort("A CLI toolbox with multiple commands"),
		gocli.WithLong("Toolbox demonstrates go-cli's subcommand capabilities."),
	)

	versionCmd := gocli.NewCommand(
		gocli.WithName("version"),
		gocli.WithAlias("v", "ver"),
		gocli.WithShort("Print version information"),
		gocli.WithRun(func(cmd *gocli.Command, args []string) error {
			fmt.Println("toolbox v1.0.0")
			fmt.Println("Built with go-cli")
			return nil
		}),
	)

	echoCmd := gocli.NewCommand(
		gocli.WithName("echo"),
		gocli.WithShort("Echo the provided arguments"),
		gocli.WithArgValidator(gocli.MinimumNArgs(1)),
		gocli.WithRun(func(cmd *gocli.Command, args []string) error {
			fmt.Println(strings.Join(args, " "))
			return nil
		}),
	)

	uppercaseCmd := gocli.NewCommand(
		gocli.WithName("uppercase"),
		gocli.WithAlias("upper", "up"),
		gocli.WithShort("Convert text to uppercase"),
		gocli.WithArgValidator(gocli.MinimumNArgs(1)),
		gocli.WithRun(func(cmd *gocli.Command, args []string) error {
			text := strings.Join(args, " ")
			fmt.Println(strings.ToUpper(text))
			return nil
		}),
	)

	lowercaseCmd := gocli.NewCommand(
		gocli.WithName("lowercase"),
		gocli.WithAlias("lower", "low"),
		gocli.WithShort("Convert text to lowercase"),
		gocli.WithArgValidator(gocli.MinimumNArgs(1)),
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
