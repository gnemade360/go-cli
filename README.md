# go-cli

A Command domain entity library for Go, inspired by Cobra but designed to integrate seamlessly with the [go-config](https://github.com/gnemade360/go-config) ecosystem.

## Features

- **Command Hierarchy**: Support for subcommands and nested command structures
- **Lifecycle Hooks**: PreRun, Run, and PostRun hooks for flexible command execution
- **Argument Validation**: Built-in validators (ExactArgs, MinimumNArgs, MaximumNArgs, RangeArgs, OnlyValidArgs, MatchAll)
- **go-config Integration**: Seamless integration with go-config for configuration management
- **Functional Options Pattern**: Flexible command construction following Go best practices
- **Context Support**: Full context.Context integration for cancellation and timeouts

## Installation

```bash
go get github.com/gnemade360/go-cli
```

## Quick Start

### Simple Command

```go
package main

import (
    "fmt"
    "github.com/gnemade360/go-cli"
)

func main() {
    rootCmd := gocli.NewCommand(
        gocli.WithUse("greet"),
        gocli.WithShort("Greet someone"),
        gocli.WithArgs(gocli.ExactArgs(1)),
        gocli.WithRun(func(cmd *gocli.Command, args []string) error {
            fmt.Printf("Hello, %s!\n", args[0])
            return nil
        }),
    )

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
```

### Command with Subcommands

```go
package main

import (
    "fmt"
    "github.com/gnemade360/go-cli"
)

func main() {
    rootCmd := gocli.NewCommand(
        gocli.WithUse("myapp"),
        gocli.WithShort("My awesome application"),
    )

    versionCmd := gocli.NewCommand(
        gocli.WithUse("version"),
        gocli.WithShort("Print version information"),
        gocli.WithRun(func(cmd *gocli.Command, args []string) error {
            fmt.Println("v1.0.0")
            return nil
        }),
    )

    rootCmd.AddCommand(versionCmd)

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
```

### Integration with go-config

```go
package main

import (
    "fmt"
    "github.com/gnemade360/go-cli"
    "github.com/gnemade360/go-config/providers/sequential"
    "github.com/gnemade360/go-config/providers/env"
    "github.com/gnemade360/go-config/providers/file"
    "github.com/gnemade360/go-config/configutil"
)

func main() {
    // Create config provider
    provider := sequential.New(
        env.New(env.WithPrefix("MYAPP")),
        file.New(file.WithPath("config.yaml")),
    )

    rootCmd := gocli.NewCommand(
        gocli.WithUse("myapp"),
        gocli.WithShort("My awesome application"),
        gocli.WithConfigProvider(provider),
        gocli.WithRun(func(cmd *gocli.Command, args []string) error {
            cfg := cmd.Config()
            host := configutil.GetString(cfg, "database.host", "localhost")
            port := configutil.GetInt(cfg, "database.port", 5432)

            fmt.Printf("Connecting to %s:%d\n", host, port)
            return nil
        }),
    )

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
```

### Lifecycle Hooks

```go
package main

import (
    "fmt"
    "github.com/gnemade360/go-cli"
)

func main() {
    rootCmd := gocli.NewCommand(
        gocli.WithUse("myapp"),
        gocli.WithShort("Application with lifecycle hooks"),

        gocli.WithPreRun(func(cmd *gocli.Command, args []string) error {
            fmt.Println("Initializing...")
            return nil
        }),

        gocli.WithRun(func(cmd *gocli.Command, args []string) error {
            fmt.Println("Running main logic...")
            return nil
        }),

        gocli.WithPostRun(func(cmd *gocli.Command, args []string) error {
            fmt.Println("Cleaning up...")
            return nil
        }),
    )

    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
    }
}
```

## Argument Validators

go-cli provides several built-in argument validators:

### ExactArgs

Requires exactly N arguments:

```go
gocli.WithArgs(gocli.ExactArgs(2))
```

### MinimumNArgs

Requires at least N arguments:

```go
gocli.WithArgs(gocli.MinimumNArgs(1))
```

### MaximumNArgs

Requires at most N arguments:

```go
gocli.WithArgs(gocli.MaximumNArgs(3))
```

### RangeArgs

Requires between min and max arguments:

```go
gocli.WithArgs(gocli.RangeArgs(1, 3))
```

### OnlyValidArgs

Validates arguments against a list of valid values:

```go
gocli.NewCommand(
    gocli.WithValidArgs("start", "stop", "restart"),
    gocli.WithArgs(gocli.OnlyValidArgs()),
)
```

### MatchAll

Combines multiple validators:

```go
gocli.WithArgs(gocli.MatchAll(
    gocli.MinimumNArgs(1),
    gocli.MaximumNArgs(3),
))
```

## Command Options

### Core Options

- `WithUse(string)` - Set the command name and usage
- `WithShort(string)` - Set short description
- `WithLong(string)` - Set long description
- `WithRun(CommandFunc)` - Set the main execution function
- `WithPreRun(CommandFunc)` - Set pre-execution hook
- `WithPostRun(CommandFunc)` - Set post-execution hook

### Validation Options

- `WithArgs(ArgsValidator)` - Set argument validator
- `WithValidArgs(...string)` - Set valid argument list

### Integration Options

- `WithConfigProvider(Provider)` - Integrate with go-config provider

## API Reference

### Command Struct

```go
type Command struct {
    // (unexported fields)
}
```

### CommandFunc

```go
type CommandFunc func(cmd *Command, args []string) error
```

### ArgsValidator

```go
type ArgsValidator func(cmd *Command, args []string) error
```

### Methods

- `Execute() error` - Execute the command
- `ExecuteContext(ctx context.Context) error` - Execute with context
- `AddCommand(...*Command)` - Add subcommands
- `Config() configprovider.Provider` - Get config provider (inherits from parent if not set)
- `Context() context.Context` - Get execution context
- `Use() string` - Get command name
- `Short() string` - Get short description
- `Long() string` - Get long description

## Error Types

### InvalidArgsError

Returned when the wrong number of arguments is provided:

```go
type InvalidArgsError struct {
    Expected string
    Received int
}
```

### InvalidArgError

Returned when an invalid argument value is provided:

```go
type InvalidArgError struct {
    Arg       string
    ValidArgs []string
}
```

## Design Philosophy

go-cli is designed with the following principles:

1. **Clean Separation of Concerns**: go-cli handles command execution, while go-config handles configuration reading
2. **Functional Options Pattern**: Following Go best practices for flexible API design
3. **Ecosystem Integration**: Built to work seamlessly with go-config, go-server, go-logger, and other go-* libraries
4. **Simplicity**: Simple, focused API that does one thing well

## Comparison with Cobra

| Feature | Cobra | go-cli |
|---------|-------|--------|
| Command entity | ✅ | ✅ |
| Subcommands | ✅ | ✅ |
| Lifecycle hooks | ✅ PreRun/Run/PostRun | ✅ PreRun/Run/PostRun |
| Argument validation | ✅ | ✅ |
| **Flag management** | ✅ Built-in (pflag) | Delegates to go-config |
| **Config reading** | Via Viper (separate) | ✅ Built-in (go-config) |
| Help generation | ✅ | 🔄 Coming soon |
| Completions | ✅ | 🔄 Planned |

**Key Differentiator**: go-cli integrates with the go-config ecosystem instead of using pflag + Viper.

## Roadmap

- [ ] Phase 1: Core Command Entity (✅ Complete)
- [ ] Phase 2: Advanced Features (Command hierarchy - ✅ Complete)
- [ ] Phase 3: Help & Usage Generation
- [ ] Phase 4: Shell Completions
- [ ] Phase 5: Fuzzy Matching & Suggestions
- [ ] Phase 6: Deprecation Warnings

## Contributing

Contributions are welcome! Please feel free to submit a Pull Request.

## License

Apache License 2.0

## Related Projects

- [go-config](https://github.com/gnemade360/go-config) - Configuration management library
- [go-map-navigator](https://github.com/gnemade360/go-map-navigator) - Map navigation and traversal
- [go-server](https://github.com/gnemade360/go-server) - Server utilities
- [go-logger](https://github.com/gnemade360/go-logger) - Logging functionality
