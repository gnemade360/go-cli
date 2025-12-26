package flags

import (
	"testing"
	"time"
)

func TestFlagSet_Parse_String(t *testing.T) {
	schema := FlagSchema{
		"name": {
			Type:        FlagString,
			Description: "Name",
			Default:     "default",
		},
	}

	tests := []struct {
		name     string
		args     []string
		expected string
		wantErr  bool
	}{
		{"long form", []string{"--name", "test"}, "test", false},
		{"long form with equals", []string{"--name=test"}, "test", false},
		{"default value", []string{}, "default", false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := NewFlagSet(schema)
			err := fs.Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				val, err := fs.GetString("name")
				if err != nil {
					t.Errorf("GetString() error = %v", err)
					return
				}
				if val != tt.expected {
					t.Errorf("GetString() = %v, want %v", val, tt.expected)
				}
			}
		})
	}
}

func TestFlagSet_Parse_Int(t *testing.T) {
	schema := FlagSchema{
		"port": {
			Type:        FlagInt,
			Short:       "p",
			Description: "Port",
			Default:     8080,
		},
	}

	tests := []struct {
		name     string
		args     []string
		expected int
		wantErr  bool
	}{
		{"long form", []string{"--port", "9000"}, 9000, false},
		{"short form", []string{"-p", "9000"}, 9000, false},
		{"long form with equals", []string{"--port=9000"}, 9000, false},
		{"short form with equals", []string{"-p=9000"}, 9000, false},
		{"default value", []string{}, 8080, false},
		{"invalid value", []string{"--port", "invalid"}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := NewFlagSet(schema)
			err := fs.Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				val, err := fs.GetInt("port")
				if err != nil {
					t.Errorf("GetInt() error = %v", err)
					return
				}
				if val != tt.expected {
					t.Errorf("GetInt() = %v, want %v", val, tt.expected)
				}
			}
		})
	}
}

func TestFlagSet_Parse_Bool(t *testing.T) {
	schema := FlagSchema{
		"debug": {
			Type:        FlagBool,
			Short:       "d",
			Description: "Debug mode",
			Default:     false,
		},
	}

	tests := []struct {
		name     string
		args     []string
		expected bool
		wantErr  bool
	}{
		{"long form without value", []string{"--debug"}, true, false},
		{"short form without value", []string{"-d"}, true, false},
		{"long form with true", []string{"--debug=true"}, true, false},
		{"long form with false", []string{"--debug=false"}, false, false},
		{"default value", []string{}, false, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := NewFlagSet(schema)
			err := fs.Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				val, err := fs.GetBool("debug")
				if err != nil {
					t.Errorf("GetBool() error = %v", err)
					return
				}
				if val != tt.expected {
					t.Errorf("GetBool() = %v, want %v", val, tt.expected)
				}
			}
		})
	}
}

func TestFlagSet_Parse_Duration(t *testing.T) {
	schema := FlagSchema{
		"timeout": {
			Type:        FlagDuration,
			Description: "Timeout",
			Default:     time.Second * 30,
		},
	}

	tests := []struct {
		name     string
		args     []string
		expected time.Duration
		wantErr  bool
	}{
		{"valid duration", []string{"--timeout", "1m"}, time.Minute, false},
		{"default value", []string{}, time.Second * 30, false},
		{"invalid duration", []string{"--timeout", "invalid"}, 0, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := NewFlagSet(schema)
			err := fs.Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				val, err := fs.GetDuration("timeout")
				if err != nil {
					t.Errorf("GetDuration() error = %v", err)
					return
				}
				if val != tt.expected {
					t.Errorf("GetDuration() = %v, want %v", val, tt.expected)
				}
			}
		})
	}
}

func TestFlagSet_Parse_StringSlice(t *testing.T) {
	schema := FlagSchema{
		"tags": {
			Type:        FlagStringSlice,
			Description: "Tags",
			Default:     []string{"default"},
		},
	}

	tests := []struct {
		name     string
		args     []string
		expected []string
		wantErr  bool
	}{
		{"single value", []string{"--tags", "a,b,c"}, []string{"a", "b", "c"}, false},
		{"default value", []string{}, []string{"default"}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := NewFlagSet(schema)
			err := fs.Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				val, err := fs.GetStringSlice("tags")
				if err != nil {
					t.Errorf("GetStringSlice() error = %v", err)
					return
				}
				if len(val) != len(tt.expected) {
					t.Errorf("GetStringSlice() length = %v, want %v", len(val), len(tt.expected))
					return
				}
				for i := range val {
					if val[i] != tt.expected[i] {
						t.Errorf("GetStringSlice()[%d] = %v, want %v", i, val[i], tt.expected[i])
					}
				}
			}
		})
	}
}

func TestFlagSet_Parse_IntSlice(t *testing.T) {
	schema := FlagSchema{
		"ports": {
			Type:        FlagIntSlice,
			Description: "Ports",
			Default:     []int{8080},
		},
	}

	tests := []struct {
		name     string
		args     []string
		expected []int
		wantErr  bool
	}{
		{"single value", []string{"--ports", "8080,8081,8082"}, []int{8080, 8081, 8082}, false},
		{"default value", []string{}, []int{8080}, false},
		{"invalid value", []string{"--ports", "8080,invalid"}, nil, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := NewFlagSet(schema)
			err := fs.Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !tt.wantErr {
				val, err := fs.GetIntSlice("ports")
				if err != nil {
					t.Errorf("GetIntSlice() error = %v", err)
					return
				}
				if len(val) != len(tt.expected) {
					t.Errorf("GetIntSlice() length = %v, want %v", len(val), len(tt.expected))
					return
				}
				for i := range val {
					if val[i] != tt.expected[i] {
						t.Errorf("GetIntSlice()[%d] = %v, want %v", i, val[i], tt.expected[i])
					}
				}
			}
		})
	}
}

func TestFlagSet_Parse_Required(t *testing.T) {
	schema := FlagSchema{
		"name": {
			Type:        FlagString,
			Description: "Name",
			Required:    true,
		},
	}

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"provided", []string{"--name", "test"}, false},
		{"missing", []string{}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := NewFlagSet(schema)
			err := fs.Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFlagSet_Parse_Validation(t *testing.T) {
	schema := FlagSchema{
		"port": {
			Type:        FlagInt,
			Description: "Port",
			Validate:    Range(1024, 65535),
		},
	}

	tests := []struct {
		name    string
		args    []string
		wantErr bool
	}{
		{"valid", []string{"--port", "8080"}, false},
		{"below range", []string{"--port", "100"}, true},
		{"above range", []string{"--port", "70000"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fs := NewFlagSet(schema)
			err := fs.Parse(tt.args)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parse() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFlagSet_Parse_UnknownFlag(t *testing.T) {
	schema := FlagSchema{
		"known": {
			Type:        FlagString,
			Description: "Known flag",
		},
	}

	fs := NewFlagSet(schema)
	err := fs.Parse([]string{"--unknown", "value"})
	if err == nil {
		t.Error("Parse() expected error for unknown flag, got nil")
	}
}

func TestFlagSet_Changed(t *testing.T) {
	schema := FlagSchema{
		"name": {
			Type:    FlagString,
			Default: "default",
		},
	}

	fs := NewFlagSet(schema)
	err := fs.Parse([]string{"--name", "test"})
	if err != nil {
		t.Errorf("Parse() error = %v", err)
		return
	}

	if !fs.Changed("name") {
		t.Error("Changed() = false, want true for changed flag")
	}
}

func TestFlagSet_Args(t *testing.T) {
	schema := FlagSchema{
		"flag": {
			Type: FlagString,
		},
	}

	fs := NewFlagSet(schema)
	err := fs.Parse([]string{"--flag", "value", "arg1", "arg2"})
	if err != nil {
		t.Errorf("Parse() error = %v", err)
		return
	}

	args := fs.Args()
	if len(args) != 2 {
		t.Errorf("Args() length = %v, want 2", len(args))
		return
	}
	if args[0] != "arg1" || args[1] != "arg2" {
		t.Errorf("Args() = %v, want [arg1 arg2]", args)
	}
}
