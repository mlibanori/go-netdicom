// Code generated by "stringer -type QRLevel"; DO NOT EDIT.

package netdicom

import "strconv"

func _() {
	// An "invalid array index" compiler error signifies that the constant values have changed.
	// Re-run the stringer command to generate them again.
	var x [1]struct{}
	_ = x[QRLevelPatient-0]
	_ = x[QRLevelStudy-1]
	_ = x[QRLevelSeries-2]
}

const _QRLevel_name = "QRLevelPatientQRLevelStudyQRLevelSeries"

var _QRLevel_index = [...]uint8{0, 14, 26, 39}

func (i QRLevel) String() string {
	if i < 0 || i >= QRLevel(len(_QRLevel_index)-1) {
		return "QRLevel(" + strconv.FormatInt(int64(i), 10) + ")"
	}
	return _QRLevel_name[_QRLevel_index[i]:_QRLevel_index[i+1]]
}
