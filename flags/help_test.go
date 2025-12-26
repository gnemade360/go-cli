package flags

import (
	"strings"
	"testing"
)

func TestHelpFormatter_GenerateHelp(t *testing.T) {
	schema := FlagSchema{
		"port": {
			Type:        FlagInt,
			Short:       "p",
			Description: "Port to listen on",
			Default:     8080,
		},
		"host": {
			Type:        FlagString,
			Short:       "h",
			Description: "Host to bind to",
			Required:    true,
		},
		"debug": {
			Type:        FlagBool,
			Description: "Enable debug mode",
			Default:     false,
		},
		"hidden": {
			Type:        FlagString,
			Description: "Hidden flag",
			Hidden:      true,
		},
		"deprecated": {
			Type:        FlagString,
			Description: "Old flag",
			Deprecated:  "use --new instead",
		},
	}

	formatter := NewHelpFormatter(schema)
	help := formatter.GenerateHelp()

	if !strings.Contains(help, "Flags:") {
		t.Error("Help should contain 'Flags:' header")
	}

	if !strings.Contains(help, "-p, --port") {
		t.Error("Help should contain port flag with short form")
	}

	if !strings.Contains(help, "Port to listen on") {
		t.Error("Help should contain port description")
	}

	if !strings.Contains(help, "(required)") {
		t.Error("Help should indicate required flag")
	}

	if !strings.Contains(help, "(default: 8080)") {
		t.Error("Help should show default value for optional flags")
	}

	if strings.Contains(help, "hidden") {
		t.Error("Help should not contain hidden flags")
	}

	if !strings.Contains(help, "DEPRECATED") {
		t.Error("Help should indicate deprecated flags")
	}
}

func TestHelpFormatter_GenerateUsage(t *testing.T) {
	schema := FlagSchema{
		"required": {
			Type:     FlagString,
			Required: true,
		},
		"optional": {
			Type: FlagString,
		},
	}

	formatter := NewHelpFormatter(schema)
	usage := formatter.GenerateUsage("mycommand")

	if !strings.Contains(usage, "Usage: mycommand") {
		t.Error("Usage should contain command name")
	}

	if !strings.Contains(usage, "[required flags]") {
		t.Error("Usage should indicate required flags")
	}

	if !strings.Contains(usage, "[flags]") {
		t.Error("Usage should indicate optional flags")
	}

	if !strings.Contains(usage, "[args...]") {
		t.Error("Usage should indicate args")
	}
}

func TestHelpFormatter_EmptySchema(t *testing.T) {
	schema := FlagSchema{}
	formatter := NewHelpFormatter(schema)
	help := formatter.GenerateHelp()

	if help != "" {
		t.Error("Empty schema should produce empty help")
	}
}

func TestFormatHelp(t *testing.T) {
	schema := FlagSchema{
		"test": {
			Type:        FlagString,
			Description: "Test flag",
		},
	}

	help := FormatHelp(schema)

	if !strings.Contains(help, "--test") {
		t.Error("FormatHelp should contain flag name")
	}
}

func TestFormatUsage(t *testing.T) {
	schema := FlagSchema{
		"test": {
			Type: FlagString,
		},
	}

	usage := FormatUsage(schema, "cmd")

	if !strings.Contains(usage, "Usage: cmd") {
		t.Error("FormatUsage should contain command name")
	}
}
