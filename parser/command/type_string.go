// Code generated by "stringer -type=Type"; DO NOT EDIT.

package command

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Arithmetic-0]
	_ = x[Push-1]
	_ = x[Pop-2]
	_ = x[Label-3]
	_ = x[Goto-4]
	_ = x[If-5]
	_ = x[Function-6]
	_ = x[Return-7]
	_ = x[Call-8]
	_ = x[Unknown-9]
}

const _Type_name = "ArithmeticPushPopLabelGotoIfFunctionReturnCallUnknown"

var _Type_index = [...]uint8{0, 10, 14, 17, 22, 26, 28, 36, 42, 46, 53}

func (i Type) String() string {
	if i < 0 || i >= Type(len(_Type_index)-1) {
		return "Type(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Type_name[_Type_index[i]:_Type_index[i+1]]
}
