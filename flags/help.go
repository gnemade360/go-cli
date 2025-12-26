package flags

import (
	"fmt"
	"sort"
	"strings"
)

type HelpFormatter struct {
	schema FlagSchema
}

func NewHelpFormatter(schema FlagSchema) *HelpFormatter {
	return &HelpFormatter{schema: schema}
}

func (h *HelpFormatter) GenerateHelp() string {
	if len(h.schema) == 0 {
		return ""
	}

	var builder strings.Builder
	builder.WriteString("Flags:\n")

	flags := h.getSortedFlags()

	for _, name := range flags {
		def := h.schema[name]
		if def.Hidden {
			continue
		}

		builder.WriteString(h.formatFlag(name, def))
	}

	return builder.String()
}

func (h *HelpFormatter) formatFlag(name string, def *FlagDefinition) string {
	var parts []string

	if def.Short != "" {
		parts = append(parts, fmt.Sprintf("-%s,", def.Short))
	}

	parts = append(parts, fmt.Sprintf("--%s", name))

	flagLine := fmt.Sprintf("  %-30s", strings.Join(parts, " "))

	var description string
	if def.Deprecated != "" {
		description = fmt.Sprintf("[DEPRECATED: %s] %s", def.Deprecated, def.Description)
	} else {
		description = def.Description
	}

	if def.Required {
		description = description + " (required)"
	}

	if def.Default != nil && !def.Required {
		description = fmt.Sprintf("%s (default: %v)", description, def.Default)
	}

	return fmt.Sprintf("%s%s\n", flagLine, description)
}

func (h *HelpFormatter) getSortedFlags() []string {
	names := make([]string, 0, len(h.schema))
	for name := range h.schema {
		names = append(names, name)
	}
	sort.Strings(names)
	return names
}

func (h *HelpFormatter) GenerateUsage(commandName string) string {
	var builder strings.Builder

	builder.WriteString(fmt.Sprintf("Usage: %s", commandName))

	flags := h.getSortedFlags()
	hasRequired := false
	hasOptional := false

	for _, name := range flags {
		def := h.schema[name]
		if def.Hidden {
			continue
		}
		if def.Required {
			hasRequired = true
		} else {
			hasOptional = true
		}
	}

	if hasRequired {
		builder.WriteString(" [required flags]")
	}
	if hasOptional {
		builder.WriteString(" [flags]")
	}

	builder.WriteString(" [args...]\n\n")

	return builder.String()
}

func FormatHelp(schema FlagSchema) string {
	formatter := NewHelpFormatter(schema)
	return formatter.GenerateHelp()
}

func FormatUsage(schema FlagSchema, commandName string) string {
	formatter := NewHelpFormatter(schema)
	return formatter.GenerateUsage(commandName)
}
