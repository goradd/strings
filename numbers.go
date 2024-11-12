package strings

import (
	"golang.org/x/exp/constraints"
	"strconv"
	"strings"
	"unicode"
)

// ExtractNumbers returns a string with the digits contained in the given string.
func ExtractNumbers(in string) string {
	return strings.Map(func(r rune) rune {
		if unicode.IsNumber(r) {
			return r
		}
		return -1
	}, in)
}

// AtoI is a convenience script for converting a string to various types of signed integers.
// An invalid input will return zero, including if the input overflows the max size of the integer type.
func AtoI[T constraints.Integer](s string) T {
	var t T
	signed := true
	bitSize := 0
	switch any(t).(type) {
	case uint:
		signed = false
	case uint8:
		signed = false
		bitSize = 8
	case uint16:
		signed = false
		bitSize = 16
	case uint32:
		signed = false
		bitSize = 32
	case uint64:
		signed = false
		bitSize = 64
	case int8:
		bitSize = 8
	case int16:
		bitSize = 16
	case int32:
		bitSize = 32
	case int64:
		bitSize = 64
	}
	if signed {
		v, err := strconv.ParseInt(s, 10, bitSize)
		if err != nil {
			return 0
		}
		return T(v)
	} else {
		v, err := strconv.ParseUint(s, 10, 0)
		if err != nil {
			return 0
		}
		return T(v)
	}
}
