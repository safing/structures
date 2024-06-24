package dsd

import (
	"encoding/hex"
	"strings"
)

// safeFirst16Bytes return the first 16 bytes of the given data in safe form.
func safeFirst16Bytes(data []byte) string {
	if len(data) == 0 {
		return "<empty>"
	}

	return strings.TrimPrefix(
		strings.SplitN(hex.Dump(data), "\n", 2)[0],
		"00000000  ",
	)
}

// safeFirst16Chars return the first 16 characters of the given data in safe form.
func safeFirst16Chars(s string) string {
	return safeFirst16Bytes([]byte(s))
}
