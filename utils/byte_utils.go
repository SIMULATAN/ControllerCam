package utils

import "fmt"

func DumpByteSlice(b []byte) string {
	result := ""

	for _, v := range b {
		result += fmt.Sprintf("%02X ", v)
	}

	return result
}

func BytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}
