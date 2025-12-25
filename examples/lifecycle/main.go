package main

import (
	"fmt"
	"os"
	"time"

	"github.com/gnemade360/go-cli"
)

func main() {
	rootCmd := gocli.NewCommand(
		gocli.WithName("lifecycle"),
		gocli.WithShort("Demonstrates lifecycle hooks"),
		gocli.WithLong("Shows how PreRun, Run, and PostRun hooks execute in sequence."),

		gocli.WithPreRun(func(cmd *gocli.Command, args []string) error {
			fmt.Println("🔧 [PreRun] Initializing application...")
			fmt.Println("   - Checking dependencies")
			fmt.Println("   - Loading configuration")
			fmt.Println("   - Setting up connections")
			time.Sleep(500 * time.Millisecond)
			fmt.Println("   ✓ Initialization complete")
			fmt.Println()
			return nil
		}),

		gocli.WithRun(func(cmd *gocli.Command, args []string) error {
			fmt.Println("⚙️  [Run] Executing main application logic...")
			fmt.Println("   - Processing data")
			fmt.Println("   - Performing operations")
			time.Sleep(1 * time.Second)
			fmt.Println("   ✓ Operations complete")
			fmt.Println()
			return nil
		}),

		gocli.WithPostRun(func(cmd *gocli.Command, args []string) error {
			fmt.Println("🧹 [PostRun] Cleaning up...")
			fmt.Println("   - Closing connections")
			fmt.Println("   - Flushing buffers")
			fmt.Println("   - Releasing resources")
			time.Sleep(300 * time.Millisecond)
			fmt.Println("   ✓ Cleanup complete")
			fmt.Println()
			return nil
		}),
	)

	errorCmd := gocli.NewCommand(
		gocli.WithName("error"),
		gocli.WithShort("Demonstrates error handling in hooks"),
		gocli.WithPreRun(func(cmd *gocli.Command, args []string) error {
			fmt.Println("🔧 [PreRun] Starting initialization...")
			return nil
		}),
		gocli.WithRun(func(cmd *gocli.Command, args []string) error {
			fmt.Println("⚙️  [Run] Simulating an error...")
			return fmt.Errorf("something went wrong during execution")
		}),
		gocli.WithPostRun(func(cmd *gocli.Command, args []string) error {
			fmt.Println("🧹 [PostRun] This won't execute due to Run error")
			return nil
		}),
	)

	rootCmd.AddCommand(errorCmd)

	fmt.Println("=== go-cli Lifecycle Demo ===")
	fmt.Println()

	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintf(os.Stderr, "\n❌ Error: %v\n", err)
		os.Exit(1)
	}

	fmt.Println("✅ Application finished successfully!")
}
