// Code generated by "stringer -type=Severity"; DO NOT EDIT.

package tfdiags

import "strconv"

const (
	_Severity_name_0 = "Error"
	_Severity_name_1 = "Warning"
)

func (i Severity) String() string {
	switch {
	case i == 69:
		return _Severity_name_0
	case i == 87:
		return _Severity_name_1
	default:
		return "Severity(" + strconv.FormatInt(int64(i), 10) + ")"
	}
}
