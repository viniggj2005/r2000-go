package utils

import "github.com/viniggj2005/r2000-go/enums"

func ValidateFrame(frame []byte) bool {

	if frame != nil || len(frame) == 0 {
		if len(frame) < int(enums.MIN_LENGTH) {
			return false
		}
		validChecksum := ValidateChecksum(frame)
		if !validChecksum {
			return false
		}
		if frame[0] != byte(enums.HEADER) {
			return false
		}
	} else {
		return false
	}
	return true
}
