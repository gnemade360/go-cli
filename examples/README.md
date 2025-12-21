# go-cli Examples

This directory contains example applications demonstrating various features of go-cli. These examples are not part of the library itself but are available for learning and testing.

## Available Examples

### 1. Basic (`basic/`)

A simple greeting application demonstrating:
- Basic command creation
- Argument validation with `ExactArgs`
- Simple command execution

**Run:**
```bash
cd examples/basic
go run main.go Alice
# Output: Hello, Alice! Welcome to go-cli.
```

**Try invalid args:**
```bash
go run main.go
# Error: invalid number of arguments: expected 1 arg(s), received 0

go run main.go Alice Bob
# Error: invalid number of arguments: expected 1 arg(s), received 2
```

---

### 2. Subcommands (`subcommands/`)

A CLI toolbox with multiple subcommands demonstrating:
- Command hierarchy
- Multiple subcommands
- Different argument validators

**Available commands:**
- `version` - Print version information
- `echo` - Echo the provided arguments
- `uppercase` - Convert text to uppercase
- `lowercase` - Convert text to lowercase

**Run:**
```bash
cd examples/subcommands

# Version command
go run main.go version
# Output: toolbox v1.0.0

# Echo command
go run main.go echo Hello World
# Output: Hello World

# Uppercase command
go run main.go uppercase hello world
# Output: HELLO WORLD

# Lowercase command
go run main.go lowercase HELLO WORLD
# Output: hello world
```

---

### 3. With Config (`with-config/`)

Demonstrates go-cli integration with go-config for configuration management:
- Config provider injection
- Environment variable configuration
- Config inheritance in subcommands

**Available commands:**
- `config` - Show current configuration values
- `connect` - Connect using configured settings

**Run:**
```bash
cd examples/with-config

# Show default config
go run main.go config
# Output:
# Configuration:
#   Host:  localhost
#   Port:  8080
#   Debug: false

# Override with environment variables
export APP_HOST=example.com
export APP_PORT=3000
export APP_DEBUG=true
go run main.go config
# Output:
# Configuration:
#   Host:  example.com
#   Port:  3000
#   Debug: true

# Use config in subcommand
go run main.go connect
# Output: Connecting to example.com:3000...
```

---

### 4. Lifecycle (`lifecycle/`)

Demonstrates lifecycle hooks (PreRun, Run, PostRun):
- Sequential execution of hooks
- Error handling in lifecycle
- Cleanup operations

**Run:**
```bash
cd examples/lifecycle

# Normal execution
go run main.go
# Shows PreRun → Run → PostRun sequence

# Error handling
go run main.go error
# Shows PreRun → Run (error) → PostRun skipped
```

**Output:**
```
=== go-cli Lifecycle Demo ===

🔧 [PreRun] Initializing application...
   - Checking dependencies
   - Loading configuration
   - Setting up connections
   ✓ Initialization complete

⚙️  [Run] Executing main application logic...
   - Processing data
   - Performing operations
   ✓ Operations complete

🧹 [PostRun] Cleaning up...
   - Closing connections
   - Flushing buffers
   - Releasing resources
   ✓ Cleanup complete

✅ Application finished successfully!
```

---

## Running All Examples

You can run all examples at once using the provided script:

```bash
# From the examples directory
./run-all.sh
```

Or run them individually as shown above.

---

## Building Examples

Each example can be built into a standalone binary:

```bash
# Build basic example
cd examples/basic
go build -o greet

# Run the binary
./greet Alice
```

```bash
# Build toolbox example
cd examples/subcommands
go build -o toolbox

# Run the binary
./toolbox version
./toolbox uppercase hello world
```

---

## Development

All examples use a `replace` directive in their `go.mod` to point to the local go-cli library:

```go
replace github.com/gnemade360/go-cli => ../..
```

This means:
- Changes to the library are immediately available in examples
- No need to publish the library to test changes
- Examples always use the latest local code

To use the published version instead, remove the `replace` directive and run:

```bash
go get github.com/gnemade360/go-cli@latest
```

---

## Notes

- These examples are not included when you `go get github.com/gnemade360/go-cli`
- They are available in the GitHub repository for reference and testing
- Each example is a complete, standalone Go program
- Examples follow go-cli best practices

---

## Adding Your Own Example

To add a new example:

1. Create a new directory under `examples/`
2. Create `main.go` with your example code
3. Create `go.mod`:
   ```bash
   go mod init example/yourname
   ```
4. Add replace directive:
   ```go
   replace github.com/gnemade360/go-cli => ../..
   ```
5. Document it in this README

---

## Learn More

- [go-cli Documentation](../)
- [go-config Documentation](https://github.com/gnemade360/go-config)
- [API Reference](../README.md#api-reference)
