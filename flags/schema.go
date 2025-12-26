package flags

type FlagDefinition struct {
	Type        FlagType
	Short       string
	Description string
	Default     interface{}
	Required    bool
	Validate    Validator
	Deprecated  string
	Hidden      bool
	Persistent  bool
	Annotations map[string]string
}

type FlagSchema map[string]*FlagDefinition
