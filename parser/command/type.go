package command

//go:generate stringer -type=Type
type Type int

const (
	Arithmetic Type = iota
	Push
	Pop
	Label
	Goto
	If
	Function
	Return
	Call
	Unknown
)

func NewType(cmd string) Type {
	switch cmd {
	case "add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not":
		return Arithmetic
	case "push":
		return Push
	case "pop":
		return Pop
	case "label":
		return Label
	case "if-goto":
		return If
	case "goto":
		return Goto
	}

	return Unknown
}
