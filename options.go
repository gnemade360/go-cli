package gocli

import (
	"github.com/gnemade360/go-cli/flags"
	"github.com/gnemade360/go-config/configprovider"
)

type CommandOption func(*Command)

func WithName(name string) CommandOption {
	return func(c *Command) {
		c.commandName = name
	}
}

func WithAlias(aliases ...string) CommandOption {
	return func(c *Command) {
		c.aliases = aliases
	}
}

func WithShort(short string) CommandOption {
	return func(c *Command) {
		c.short = short
	}
}

func WithLong(long string) CommandOption {
	return func(c *Command) {
		c.long = long
	}
}

func WithRun(run CommandFunc) CommandOption {
	return func(c *Command) {
		c.run = run
	}
}

func WithPreRun(preRun CommandFunc) CommandOption {
	return func(c *Command) {
		c.preRun = preRun
	}
}

func WithPostRun(postRun CommandFunc) CommandOption {
	return func(c *Command) {
		c.postRun = postRun
	}
}

func WithArgValidator(validator ArgsValidator) CommandOption {
	return func(c *Command) {
		c.argValidation = validator
	}
}

func WithAllowedArgs(args ...string) CommandOption {
	return func(c *Command) {
		c.allowedArgs = args
	}
}

func WithConfigProvider(provider configprovider.Provider) CommandOption {
	return func(c *Command) {
		c.configProvider = provider
	}
}

func WithFlags(schema flags.FlagSchema) CommandOption {
	return func(c *Command) {
		c.flagSchema = schema
	}
}
