package segment

//go:generate stringer -type=Segment
type Segment int

const (
	Argument Segment = iota
	Local
	Static
	Constant
	This
	That
	Pointer
	Temp
	Unknown
)

func NewSegment(s string) Segment {
	switch s {
	case "argument":
		return Argument
	case "local":
		return Local
	case "static":
		return Static
	case "constant":
		return Constant
	case "this":
		return This
	case "that":
		return That
	case "pointer":
		return Pointer
	case "temp":
		return Temp
	}

	return Unknown
}
