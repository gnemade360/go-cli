package main

import (
	"fmt"
	"os"

	"github.com/gnemade360/go-cli"
	"github.com/gnemade360/go-config/configutil"
	"github.com/gnemade360/go-config/providers/env"
	"github.com/gnemade360/go-config/providers/sequential"
)

func main() {
	provider := sequential.New(
		sequential.WithProviders(
			env.New(env.WithPrefix("APP_")),
		),
	)

	rootCmd := gocli.NewCommand(
		gocli.WithUse("myapp"),
		gocli.WithShort("Application with configuration support"),
		gocli.WithLong("Demonstrates go-cli integration with go-config for configuration management."),
		gocli.WithConfigProvider(provider),
	)

	configCmd := gocli.NewCommand(
		gocli.WithUse("config"),
		gocli.WithShort("Show configuration values"),
		gocli.WithRun(func(cmd *gocli.Command, args []string) error {
			cfg := cmd.Config()

			host := configutil.GetString(cfg, "HOST", "localhost")
			port := configutil.GetInt(cfg, "PORT", 8080)
			debug := configutil.GetBool(cfg, "DEBUG", false)

			fmt.Println("Configuration:")
			fmt.Printf("  Host:  %s\n", host)
			fmt.Printf("  Port:  %d\n", port)
			fmt.Printf("  Debug: %t\n", debug)
			fmt.Println()
			fmt.Println("Try setting environment variables:")
			fmt.Println("  export APP_HOST=example.com")
			fmt.Println("  export APP_PORT=3000")
			fmt.Println("  export APP_DEBUG=true")
			return nil
		}),
	)

	connectCmd := gocli.NewCommand(
		gocli.WithUse("connect"),
		gocli.WithShort("Connect to configured server"),
		gocli.WithRun(func(cmd *gocli.Command, args []string) error {
			cfg := cmd.Config()

			host := configutil.GetString(cfg, "HOST", "localhost")
			port := configutil.GetInt(cfg, "PORT", 8080)

			fmt.Printf("Connecting to %s:%d...\n", host, port)
			fmt.Println("(This is a demo - not actually connecting)")
			return nil
		}),
	)

	rootCmd.AddCommand(configCmd, connectCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
