package arithmetic

//go:generate stringer -type=Command
type Command int

// "add", "sub", "neg", "eq", "gt", "lt", "and", "or", "not"
const (
	Add Command = iota
	Sub
	Neg
	EQ
	GT
	LT
	And
	Or
	Not
	Unknown
)

func NewCommand(cmd string) Command {
	switch cmd {
	case "add":
		return Add
	case "sub":
		return Sub
	case "neg":
		return Neg
	case "eq":
		return EQ
	case "gt":
		return GT
	case "lt":
		return LT
	case "and":
		return And
	case "or":
		return Or
	case "not":
		return Not
	}

	return Unknown
}

func (c Command) DoNeedTwoArgs() bool {
	switch c {
	case Neg, Not:
		return false
	}

	return true
}
