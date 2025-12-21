package gocli

import "github.com/gnemade360/go-config/configprovider"

type CommandOption func(*Command)

func WithUse(use string) CommandOption {
	return func(c *Command) {
		c.use = use
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

func WithArgs(validator ArgsValidator) CommandOption {
	return func(c *Command) {
		c.args = validator
	}
}

func WithValidArgs(validArgs ...string) CommandOption {
	return func(c *Command) {
		c.validArgs = validArgs
	}
}

func WithConfigProvider(provider configprovider.Provider) CommandOption {
	return func(c *Command) {
		c.configProvider = provider
	}
}
