package gocli

import (
	"context"
	"errors"
	"os"
	"testing"
)

func TestNewCommand(t *testing.T) {
	cmd := NewCommand(
		WithName("test"),
		WithShort("Test command"),
		WithLong("This is a test command"),
	)

	if cmd == nil {
		t.Fatal("NewCommand returned nil")
	}

	if cmd.Name() != "test" {
		t.Errorf("expected name='test', got '%s'", cmd.Name())
	}

	if cmd.Short() != "Test command" {
		t.Errorf("expected short='Test command', got '%s'", cmd.Short())
	}

	if cmd.Long() != "This is a test command" {
		t.Errorf("expected long='This is a test command', got '%s'", cmd.Long())
	}
}

func TestCommand_Execute(t *testing.T) {
	executed := false

	cmd := NewCommand(
		WithName("test"),
		WithRun(func(cmd *Command, args []string) error {
			executed = true
			return nil
		}),
	)

	oldArgs := os.Args
	os.Args = []string{"test"}
	defer func() { os.Args = oldArgs }()

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if !executed {
		t.Error("Run function was not executed")
	}
}

func TestCommand_ExecuteContext(t *testing.T) {
	var receivedCtx context.Context

	cmd := NewCommand(
		WithName("test"),
		WithRun(func(cmd *Command, args []string) error {
			receivedCtx = cmd.Context()
			return nil
		}),
	)

	oldArgs := os.Args
	os.Args = []string{"test"}
	defer func() { os.Args = oldArgs }()

	ctx := context.WithValue(context.Background(), "key", "value")
	err := cmd.ExecuteContext(ctx)
	if err != nil {
		t.Fatalf("ExecuteContext failed: %v", err)
	}

	if receivedCtx == nil {
		t.Error("Context was not set")
	}

	if receivedCtx.Value("key") != "value" {
		t.Error("Context value was not preserved")
	}
}

func TestCommand_LifecycleHooks(t *testing.T) {
	order := []string{}

	cmd := NewCommand(
		WithName("test"),
		WithPreRun(func(cmd *Command, args []string) error {
			order = append(order, "preRun")
			return nil
		}),
		WithRun(func(cmd *Command, args []string) error {
			order = append(order, "run")
			return nil
		}),
		WithPostRun(func(cmd *Command, args []string) error {
			order = append(order, "postRun")
			return nil
		}),
	)

	oldArgs := os.Args
	os.Args = []string{"test"}
	defer func() { os.Args = oldArgs }()

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	expected := []string{"preRun", "run", "postRun"}
	if len(order) != len(expected) {
		t.Fatalf("expected %d hooks, got %d", len(expected), len(order))
	}

	for i, hook := range expected {
		if order[i] != hook {
			t.Errorf("expected hook %d to be '%s', got '%s'", i, hook, order[i])
		}
	}
}

func TestCommand_LifecycleErrorHandling(t *testing.T) {
	testError := errors.New("test error")

	t.Run("PreRunError", func(t *testing.T) {
		runExecuted := false

		cmd := NewCommand(
			WithName("test"),
			WithPreRun(func(cmd *Command, args []string) error {
				return testError
			}),
			WithRun(func(cmd *Command, args []string) error {
				runExecuted = true
				return nil
			}),
		)

		oldArgs := os.Args
		os.Args = []string{"test"}
		defer func() { os.Args = oldArgs }()

		err := cmd.Execute()
		if err == nil {
			t.Error("expected error, got nil")
		}

		if runExecuted {
			t.Error("Run should not have executed after PreRun error")
		}
	})

	t.Run("RunError", func(t *testing.T) {
		postRunExecuted := false

		cmd := NewCommand(
			WithName("test"),
			WithRun(func(cmd *Command, args []string) error {
				return testError
			}),
			WithPostRun(func(cmd *Command, args []string) error {
				postRunExecuted = true
				return nil
			}),
		)

		oldArgs := os.Args
		os.Args = []string{"test"}
		defer func() { os.Args = oldArgs }()

		err := cmd.Execute()
		if err == nil {
			t.Error("expected error, got nil")
		}

		if postRunExecuted {
			t.Error("PostRun should not have executed after Run error")
		}
	})
}

func TestCommand_AddCommand(t *testing.T) {
	rootCmd := NewCommand(WithName("root"))
	subCmd := NewCommand(WithName("sub"))

	rootCmd.AddCommand(subCmd)

	if len(rootCmd.commands) != 1 {
		t.Fatalf("expected 1 subcommand, got %d", len(rootCmd.commands))
	}

	if rootCmd.commands[0] != subCmd {
		t.Error("subcommand was not added correctly")
	}

	if subCmd.parent != rootCmd {
		t.Error("parent was not set correctly")
	}
}

func TestCommand_SubcommandExecution(t *testing.T) {
	rootExecuted := false
	subExecuted := false

	rootCmd := NewCommand(
		WithName("root"),
		WithRun(func(cmd *Command, args []string) error {
			rootExecuted = true
			return nil
		}),
	)

	subCmd := NewCommand(
		WithName("sub"),
		WithRun(func(cmd *Command, args []string) error {
			subExecuted = true
			return nil
		}),
	)

	rootCmd.AddCommand(subCmd)

	oldArgs := os.Args
	os.Args = []string{"root", "sub"}
	defer func() { os.Args = oldArgs }()

	err := rootCmd.Execute()
	if err != nil {
		t.Fatalf("Execute failed: %v", err)
	}

	if rootExecuted {
		t.Error("root command should not have executed")
	}

	if !subExecuted {
		t.Error("subcommand was not executed")
	}
}

func TestCommand_ConfigInheritance(t *testing.T) {
	mockProvider := &mockConfigProvider{}

	rootCmd := NewCommand(
		WithName("root"),
		WithConfigProvider(mockProvider),
	)

	subCmd := NewCommand(WithName("sub"))
	rootCmd.AddCommand(subCmd)

	if subCmd.Config() != mockProvider {
		t.Error("subcommand did not inherit parent's config provider")
	}
}

func TestCommand_Aliases(t *testing.T) {
	cmd := NewCommand(
		WithName("version"),
		WithAlias("v", "ver"),
	)

	if len(cmd.Aliases()) != 2 {
		t.Errorf("expected 2 aliases, got %d", len(cmd.Aliases()))
	}

	if cmd.Aliases()[0] != "v" {
		t.Errorf("expected first alias 'v', got '%s'", cmd.Aliases()[0])
	}

	if cmd.Aliases()[1] != "ver" {
		t.Errorf("expected second alias 'ver', got '%s'", cmd.Aliases()[1])
	}
}

func TestCommand_AliasExecution(t *testing.T) {
	executed := false
	executedCmd := ""

	rootCmd := NewCommand(
		WithName("toolbox"),
	)

	versionCmd := NewCommand(
		WithName("version"),
		WithAlias("v", "ver"),
		WithRun(func(cmd *Command, args []string) error {
			executed = true
			executedCmd = cmd.Name()
			return nil
		}),
	)

	rootCmd.AddCommand(versionCmd)

	oldArgs := os.Args

	t.Run("execute with primary name", func(t *testing.T) {
		executed = false
		os.Args = []string{"toolbox", "version"}
		defer func() { os.Args = oldArgs }()

		err := rootCmd.Execute()
		if err != nil {
			t.Fatalf("Execute failed: %v", err)
		}

		if !executed {
			t.Error("command should have executed")
		}

		if executedCmd != "version" {
			t.Errorf("expected command name 'version', got '%s'", executedCmd)
		}
	})

	t.Run("execute with first alias", func(t *testing.T) {
		executed = false
		os.Args = []string{"toolbox", "v"}
		defer func() { os.Args = oldArgs }()

		err := rootCmd.Execute()
		if err != nil {
			t.Fatalf("Execute failed: %v", err)
		}

		if !executed {
			t.Error("command should have executed via alias 'v'")
		}

		if executedCmd != "version" {
			t.Errorf("expected command name 'version', got '%s'", executedCmd)
		}
	})

	t.Run("execute with second alias", func(t *testing.T) {
		executed = false
		os.Args = []string{"toolbox", "ver"}
		defer func() { os.Args = oldArgs }()

		err := rootCmd.Execute()
		if err != nil {
			t.Fatalf("Execute failed: %v", err)
		}

		if !executed {
			t.Error("command should have executed via alias 'ver'")
		}

		if executedCmd != "version" {
			t.Errorf("expected command name 'version', got '%s'", executedCmd)
		}
	})
}

type mockConfigProvider struct{}

func (m *mockConfigProvider) Read(key string) (interface{}, error) {
	return nil, nil
}
