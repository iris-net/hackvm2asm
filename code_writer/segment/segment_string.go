// Code generated by "stringer -type=Segment"; DO NOT EDIT.

package segment

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Argument-0]
	_ = x[Local-1]
	_ = x[Static-2]
	_ = x[Constant-3]
	_ = x[This-4]
	_ = x[That-5]
	_ = x[Pointer-6]
	_ = x[Temp-7]
	_ = x[Unknown-8]
}

const _Segment_name = "ArgumentLocalStaticConstantThisThatPointerTempUnknown"

var _Segment_index = [...]uint8{0, 8, 13, 19, 27, 31, 35, 42, 46, 53}

func (i Segment) String() string {
	if i < 0 || i >= Segment(len(_Segment_index)-1) {
		return "Segment(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Segment_name[_Segment_index[i]:_Segment_index[i+1]]
}
