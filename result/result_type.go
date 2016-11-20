package result

// Type represents result type
type Type int

const (
	// TopResultType - results sorted by top
	TopResultType Type = 1 << iota

	// BottomResultType - results sorted by bottom
	BottomResultType
)

// you could use stringer to get type names
