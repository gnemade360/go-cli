package flags

type FlagType int

const (
	FlagString FlagType = iota
	FlagInt
	FlagInt64
	FlagFloat64
	FlagBool
	FlagDuration
	FlagStringSlice
	FlagIntSlice
)

func (t FlagType) String() string {
	switch t {
	case FlagString:
		return "string"
	case FlagInt:
		return "int"
	case FlagInt64:
		return "int64"
	case FlagFloat64:
		return "float64"
	case FlagBool:
		return "bool"
	case FlagDuration:
		return "duration"
	case FlagStringSlice:
		return "[]string"
	case FlagIntSlice:
		return "[]int"
	default:
		return "unknown"
	}
}
