package flags

import (
	"testing"
	"time"
)

func TestRange(t *testing.T) {
	validator := Range(10, 20)

	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"valid", 15, false},
		{"min", 10, false},
		{"max", 20, false},
		{"below min", 5, true},
		{"above max", 25, true},
		{"wrong type", "string", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Range() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestPositive(t *testing.T) {
	validator := Positive()

	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"positive", 5, false},
		{"zero", 0, true},
		{"negative", -5, true},
		{"wrong type", "string", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Positive() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNotEmpty(t *testing.T) {
	validator := NotEmpty()

	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"non-empty", "test", false},
		{"empty", "", true},
		{"whitespace only", "   ", true},
		{"wrong type", 123, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("NotEmpty() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestOneOf(t *testing.T) {
	validator := OneOf("dev", "staging", "prod")

	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"valid dev", "dev", false},
		{"valid staging", "staging", false},
		{"valid prod", "prod", false},
		{"invalid", "test", true},
		{"wrong type", 123, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("OneOf() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRegex(t *testing.T) {
	validator := Regex(`^[a-z]+$`)

	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"valid", "test", false},
		{"invalid uppercase", "Test", true},
		{"invalid numbers", "test123", true},
		{"wrong type", 123, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("Regex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMinLength(t *testing.T) {
	validator := MinLength(5)

	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"valid", "12345", false},
		{"longer", "123456", false},
		{"too short", "1234", true},
		{"wrong type", 123, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMaxLength(t *testing.T) {
	validator := MaxLength(5)

	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"valid", "12345", false},
		{"shorter", "1234", false},
		{"too long", "123456", true},
		{"wrong type", 123, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("MaxLength() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMinDuration(t *testing.T) {
	validator := MinDuration(time.Second * 10)

	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"valid", time.Second * 15, false},
		{"exact", time.Second * 10, false},
		{"too short", time.Second * 5, true},
		{"wrong type", "string", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinDuration() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMinItems(t *testing.T) {
	validator := MinItems(2)

	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"valid string slice", []string{"a", "b"}, false},
		{"valid int slice", []int{1, 2}, false},
		{"too few string", []string{"a"}, true},
		{"too few int", []int{1}, true},
		{"wrong type", "string", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("MinItems() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestUniqueItems(t *testing.T) {
	validator := UniqueItems()

	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"unique", []string{"a", "b", "c"}, false},
		{"duplicate", []string{"a", "b", "a"}, true},
		{"wrong type", 123, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("UniqueItems() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestAll(t *testing.T) {
	validator := All(
		MinLength(5),
		MaxLength(10),
		Regex(`^[a-z]+$`),
	)

	tests := []struct {
		name    string
		value   interface{}
		wantErr bool
	}{
		{"valid", "hello", false},
		{"too short", "hi", true},
		{"too long", "verylongstring", true},
		{"invalid pattern", "Hello", true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := validator(tt.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("All() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
