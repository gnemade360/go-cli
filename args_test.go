package gocli

import (
	"os"
	"testing"
)

func TestExactArgs(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		args      []string
		wantError bool
	}{
		{"exact match", 2, []string{"a", "b"}, false},
		{"too few", 2, []string{"a"}, true},
		{"too many", 2, []string{"a", "b", "c"}, true},
		{"zero args match", 0, []string{}, false},
		{"zero args mismatch", 0, []string{"a"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewCommand(
				WithName("test"),
				WithArgValidator(ExactArgs(tt.n)),
			)

			err := cmd.argValidation(cmd, tt.args)
			if (err != nil) != tt.wantError {
				t.Errorf("ExactArgs(%d) with %d args: error = %v, wantError %v",
					tt.n, len(tt.args), err, tt.wantError)
			}
		})
	}
}

func TestMinimumNArgs(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		args      []string
		wantError bool
	}{
		{"exact match", 2, []string{"a", "b"}, false},
		{"more than minimum", 2, []string{"a", "b", "c"}, false},
		{"less than minimum", 2, []string{"a"}, true},
		{"zero minimum", 0, []string{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewCommand(
				WithName("test"),
				WithArgValidator(MinimumNArgs(tt.n)),
			)

			err := cmd.argValidation(cmd, tt.args)
			if (err != nil) != tt.wantError {
				t.Errorf("MinimumNArgs(%d) with %d args: error = %v, wantError %v",
					tt.n, len(tt.args), err, tt.wantError)
			}
		})
	}
}

func TestMaximumNArgs(t *testing.T) {
	tests := []struct {
		name      string
		n         int
		args      []string
		wantError bool
	}{
		{"exact match", 2, []string{"a", "b"}, false},
		{"less than maximum", 2, []string{"a"}, false},
		{"more than maximum", 2, []string{"a", "b", "c"}, true},
		{"zero maximum with args", 0, []string{"a"}, true},
		{"zero maximum no args", 0, []string{}, false},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewCommand(
				WithName("test"),
				WithArgValidator(MaximumNArgs(tt.n)),
			)

			err := cmd.argValidation(cmd, tt.args)
			if (err != nil) != tt.wantError {
				t.Errorf("MaximumNArgs(%d) with %d args: error = %v, wantError %v",
					tt.n, len(tt.args), err, tt.wantError)
			}
		})
	}
}

func TestRangeArgs(t *testing.T) {
	tests := []struct {
		name      string
		min       int
		max       int
		args      []string
		wantError bool
	}{
		{"within range", 1, 3, []string{"a", "b"}, false},
		{"at minimum", 1, 3, []string{"a"}, false},
		{"at maximum", 1, 3, []string{"a", "b", "c"}, false},
		{"below minimum", 1, 3, []string{}, true},
		{"above maximum", 1, 3, []string{"a", "b", "c", "d"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewCommand(
				WithName("test"),
				WithArgValidator(RangeArgs(tt.min, tt.max)),
			)

			err := cmd.argValidation(cmd, tt.args)
			if (err != nil) != tt.wantError {
				t.Errorf("RangeArgs(%d, %d) with %d args: error = %v, wantError %v",
					tt.min, tt.max, len(tt.args), err, tt.wantError)
			}
		})
	}
}

func TestOnlyValidArgs(t *testing.T) {
	tests := []struct {
		name      string
		validArgs []string
		args      []string
		wantError bool
	}{
		{"all valid", []string{"start", "stop"}, []string{"start"}, false},
		{"multiple valid", []string{"start", "stop", "restart"}, []string{"start", "stop"}, false},
		{"one invalid", []string{"start", "stop"}, []string{"start", "pause"}, true},
		{"all invalid", []string{"start", "stop"}, []string{"pause", "resume"}, true},
		{"empty valid args", []string{}, []string{"anything"}, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			cmd := NewCommand(
				WithName("test"),
				WithAllowedArgs(tt.validArgs...),
				WithArgValidator(OnlyValidArgs()),
			)

			err := cmd.argValidation(cmd, tt.args)
			if (err != nil) != tt.wantError {
				t.Errorf("OnlyValidArgs() with args %v: error = %v, wantError %v",
					tt.args, err, tt.wantError)
			}
		})
	}
}

func TestMatchAll(t *testing.T) {
	t.Run("all validators pass", func(t *testing.T) {
		cmd := NewCommand(
			WithName("test"),
			WithArgValidator(MatchAll(
				MinimumNArgs(1),
				MaximumNArgs(3),
			)),
		)

		err := cmd.argValidation(cmd, []string{"a", "b"})
		if err != nil {
			t.Errorf("MatchAll should pass, got error: %v", err)
		}
	})

	t.Run("first validator fails", func(t *testing.T) {
		cmd := NewCommand(
			WithName("test"),
			WithArgValidator(MatchAll(
				MinimumNArgs(3),
				MaximumNArgs(5),
			)),
		)

		err := cmd.argValidation(cmd, []string{"a"})
		if err == nil {
			t.Error("MatchAll should fail on first validator")
		}
	})

	t.Run("second validator fails", func(t *testing.T) {
		cmd := NewCommand(
			WithName("test"),
			WithArgValidator(MatchAll(
				MinimumNArgs(1),
				MaximumNArgs(2),
			)),
		)

		err := cmd.argValidation(cmd, []string{"a", "b", "c"})
		if err == nil {
			t.Error("MatchAll should fail on second validator")
		}
	})
}

func TestArgsValidation_Integration(t *testing.T) {
	executed := false

	cmd := NewCommand(
		WithName("test"),
		WithArgValidator(ExactArgs(2)),
		WithRun(func(cmd *Command, args []string) error {
			executed = true
			return nil
		}),
	)

	oldArgs := os.Args

	t.Run("valid args count", func(t *testing.T) {
		executed = false
		os.Args = []string{"test", "arg1", "arg2"}
		defer func() { os.Args = oldArgs }()

		err := cmd.Execute()
		if err != nil {
			t.Fatalf("Execute with valid args failed: %v", err)
		}

		if !executed {
			t.Error("command should have executed")
		}
	})

	t.Run("invalid args count", func(t *testing.T) {
		executed = false
		os.Args = []string{"test", "arg1"}
		defer func() { os.Args = oldArgs }()

		err := cmd.Execute()
		if err == nil {
			t.Error("Execute should fail with invalid arg count")
		}

		if executed {
			t.Error("command should not have executed")
		}
	})
}
