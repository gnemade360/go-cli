package main

import (
	"fmt"
	"os"
	"time"

	gocli "github.com/gnemade360/go-cli"
	"github.com/gnemade360/go-cli/flags"
)

func main() {
	rootCmd := gocli.NewCommand(
		gocli.WithName("server"),
		gocli.WithShort("HTTP server with configurable options"),
		gocli.WithLong("A sample HTTP server demonstrating flag handling with validation and type safety"),
		gocli.WithFlags(flags.FlagSchema{
			"port": {
				Type:        flags.FlagInt,
				Short:       "p",
				Description: "Port to listen on",
				Default:     8080,
				Validate:    flags.Range(1024, 65535),
			},
			"host": {
				Type:        flags.FlagString,
				Short:       "h",
				Description: "Host to bind to",
				Default:     "localhost",
				Validate:    flags.NotEmpty(),
			},
			"debug": {
				Type:        flags.FlagBool,
				Short:       "d",
				Description: "Enable debug mode",
				Default:     false,
			},
			"timeout": {
				Type:        flags.FlagDuration,
				Short:       "t",
				Description: "Request timeout",
				Default:     30 * time.Second,
			},
			"env": {
				Type:        flags.FlagString,
				Short:       "e",
				Description: "Environment",
				Default:     "development",
				Validate:    flags.OneOf("development", "staging", "production"),
			},
			"allowed-origins": {
				Type:        flags.FlagStringSlice,
				Description: "Allowed CORS origins (comma-separated)",
				Default:     []string{"http://localhost:3000"},
			},
			"max-connections": {
				Type:        flags.FlagInt,
				Description: "Maximum concurrent connections",
				Default:     100,
				Validate:    flags.Positive(),
			},
		}),
		gocli.WithRun(runServer),
	)

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

func runServer(cmd *gocli.Command, args []string) error {
	port, err := cmd.Flags().GetInt("port")
	if err != nil {
		return err
	}

	host, err := cmd.Flags().GetString("host")
	if err != nil {
		return err
	}

	debug, err := cmd.Flags().GetBool("debug")
	if err != nil {
		return err
	}

	timeout, err := cmd.Flags().GetDuration("timeout")
	if err != nil {
		return err
	}

	env, err := cmd.Flags().GetString("env")
	if err != nil {
		return err
	}

	origins, err := cmd.Flags().GetStringSlice("allowed-origins")
	if err != nil {
		return err
	}

	maxConn, err := cmd.Flags().GetInt("max-connections")
	if err != nil {
		return err
	}

	fmt.Println("Server Configuration:")
	fmt.Printf("  Address:            %s:%d\n", host, port)
	fmt.Printf("  Environment:        %s\n", env)
	fmt.Printf("  Debug Mode:         %v\n", debug)
	fmt.Printf("  Request Timeout:    %v\n", timeout)
	fmt.Printf("  Max Connections:    %d\n", maxConn)
	fmt.Printf("  Allowed Origins:    %v\n", origins)
	fmt.Println()
	fmt.Println("Server would start here...")

	return nil
}
