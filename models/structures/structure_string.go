// Code generated by "stringer -type Structure -linecomment"; DO NOT EDIT.

package structures

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Unknown-0]
	_ = x[AtGrade-1]
	_ = x[Elevated-2]
	_ = x[Embankment-3]
	_ = x[OpenCut-4]
	_ = x[Subway-5]
	_ = x[Viaduct-6]
}

const _Structure_name = "UnknownAt GradeElevatedEmbankmentOpen CutSubwayViaduct"

var _Structure_index = [...]uint8{0, 7, 15, 23, 33, 41, 47, 54}

func (i Structure) String() string {
	if i >= Structure(len(_Structure_index)-1) {
		return "Structure(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Structure_name[_Structure_index[i]:_Structure_index[i+1]]
}
