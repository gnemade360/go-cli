package gocli

import (
	"context"
	"fmt"
	"os"

	"github.com/gnemade360/go-cli/flags"
	"github.com/gnemade360/go-config/configprovider"
)

type Command struct {
	commandName string
	aliases     []string
	short       string
	long        string

	parent   *Command
	commands []*Command

	preRun  CommandFunc
	run     CommandFunc
	postRun CommandFunc

	argValidation ArgsValidator
	allowedArgs   []string

	flagSchema flags.FlagSchema
	flagSet    *flags.FlagSet

	configProvider configprovider.Provider

	ctx context.Context
}

type CommandFunc func(cmd *Command, args []string) error

type ArgsValidator func(cmd *Command, args []string) error

func NewCommand(opts ...CommandOption) *Command {
	cmd := &Command{
		commands: make([]*Command, 0),
		ctx:      context.Background(),
	}

	for _, opt := range opts {
		opt(cmd)
	}

	return cmd
}

func (c *Command) AddCommand(commands ...*Command) {
	for _, cmd := range commands {
		cmd.parent = c
		if cmd.configProvider == nil && c.configProvider != nil {
			cmd.configProvider = c.configProvider
		}
		c.commands = append(c.commands, cmd)
	}
}

func (c *Command) Execute() error {
	return c.ExecuteContext(context.Background())
}

func (c *Command) ExecuteContext(ctx context.Context) error {
	c.ctx = ctx

	args := os.Args[1:]

	target, targetArgs, err := c.findTarget(args)
	if err != nil {
		return err
	}

	if target.flagSchema != nil {
		target.flagSet = flags.NewFlagSet(target.flagSchema)
		if err := target.flagSet.Parse(targetArgs); err != nil {
			return err
		}
		targetArgs = target.flagSet.Args()
	}

	if target.argValidation != nil {
		if err := target.argValidation(target, targetArgs); err != nil {
			return err
		}
	}

	return target.executeLifecycle(targetArgs)
}

func (c *Command) executeLifecycle(args []string) error {
	if c.preRun != nil {
		if err := c.preRun(c, args); err != nil {
			return fmt.Errorf("preRun failed: %w", err)
		}
	}

	if c.run != nil {
		if err := c.run(c, args); err != nil {
			return fmt.Errorf("run failed: %w", err)
		}
	}

	if c.postRun != nil {
		if err := c.postRun(c, args); err != nil {
			return fmt.Errorf("postRun failed: %w", err)
		}
	}

	return nil
}

func (c *Command) findTarget(args []string) (*Command, []string, error) {
	if len(args) == 0 {
		return c, args, nil
	}

	for _, cmd := range c.commands {
		if cmd.commandName == args[0] {
			return cmd.findTarget(args[1:])
		}

		// Check aliases
		for _, alias := range cmd.aliases {
			if alias == args[0] {
				return cmd.findTarget(args[1:])
			}
		}
	}

	return c, args, nil
}

func (c *Command) Config() configprovider.Provider {
	if c.configProvider != nil {
		return c.configProvider
	}

	if c.parent != nil {
		return c.parent.Config()
	}

	return nil
}

func (c *Command) Context() context.Context {
	return c.ctx
}

func (c *Command) Name() string {
	return c.commandName
}

func (c *Command) Aliases() []string {
	return c.aliases
}

func (c *Command) Short() string {
	return c.short
}

func (c *Command) Long() string {
	return c.long
}

func (c *Command) Flags() *flags.FlagSet {
	return c.flagSet
}
