// Code generated by "stringer -type Route -linecomment"; DO NOT EDIT.

package routes

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[Line1-0]
	_ = x[Line2-1]
	_ = x[Line3-2]
	_ = x[Line4-3]
	_ = x[Line5-4]
	_ = x[Line6-5]
	_ = x[Line7-6]
	_ = x[LineA-7]
	_ = x[LineB-8]
	_ = x[LineC-9]
	_ = x[LineD-10]
	_ = x[LineE-11]
	_ = x[LineF-12]
	_ = x[LineG-13]
	_ = x[LineJ-14]
	_ = x[LineL-15]
	_ = x[LineM-16]
	_ = x[LineN-17]
	_ = x[LineQ-18]
	_ = x[LineR-19]
	_ = x[LineS-20]
	_ = x[SIR-21]
	_ = x[LineW-22]
	_ = x[LineZ-23]
	_ = x[UnknownRoute-24]
}

const _Route_name = "1234567ABCDEFGJLMNQRSSIRWZUnknown"

var _Route_index = [...]uint8{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 24, 25, 26, 33}

func (i Route) String() string {
	if i >= Route(len(_Route_index)-1) {
		return "Route(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _Route_name[_Route_index[i]:_Route_index[i+1]]
}
