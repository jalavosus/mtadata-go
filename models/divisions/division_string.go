// Code generated by "stringer -type Division -linecomment"; DO NOT EDIT.

package divisions

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[BMT-1]
	_ = x[IND-2]
	_ = x[IRT-3]
	_ = x[SIR-4]
}

const _Division_name = "UnknownBMTINDIRTSIR"

var _Division_index = [...]uint8{0, 7, 10, 13, 16, 19}

func (i Division) String() string {
	if i >= Division(len(_Division_index)-1) {
		return "Division(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Division_name[_Division_index[i]:_Division_index[i+1]]
}