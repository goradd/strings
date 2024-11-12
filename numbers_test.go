package strings

import (
	"fmt"
	"golang.org/x/exp/constraints"
	"math"
	"strconv"
	"testing"
)

func ExampleExtractNumbers() {
	a := ExtractNumbers("a1b2 c3")
	fmt.Println(a)
	//Output: 123
}

func TestExtractNumbers(t *testing.T) {
	tests := []struct {
		name    string
		in      string
		wantOut string
	}{
		{"empty", "", ""},
		{"123", "123", "123"},
		{"abc", "abc", ""},
		{"a1c", "a1c", "1"},
		{"a1c", "a1c", "1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotOut := ExtractNumbers(tt.in); gotOut != tt.wantOut {
				t.Errorf("ExtractNumbers() = %v, want %v", gotOut, tt.wantOut)
			}
		})
	}
}

// TestAtoI tests the AtoI function for various integer types.
func TestAtoI(t *testing.T) {
	type testCase[T constraints.Integer] struct {
		input    string
		expected T
	}

	// Test cases for int8
	testCasesInt8 := []testCase[int8]{
		{"0", 0},
		{"127", 127},
		{"-128", -128},
		{"128", 0},   // Overflow
		{"-129", 0},  // Underflow
		{"abc", 0},   // Invalid input
		{"12abc", 0}, // Invalid input
	}

	for _, tc := range testCasesInt8 {
		result := AtoI[int8](tc.input)
		if result != tc.expected {
			t.Errorf("AtoI[int8](%q) = %d; want %d", tc.input, result, tc.expected)
		}
	}

	// Test cases for uint8
	testCasesUint8 := []testCase[uint8]{
		{"0", 0},
		{"255", 255},
		{"256", 0},   // Overflow
		{"-1", 0},    // Negative number for unsigned
		{"abc", 0},   // Invalid input
		{"12abc", 0}, // Invalid input
	}

	for _, tc := range testCasesUint8 {
		result := AtoI[uint8](tc.input)
		if result != tc.expected {
			t.Errorf("AtoI[uint8](%q) = %d; want %d", tc.input, result, tc.expected)
		}
	}

	// Test cases for int16
	testCasesInt16 := []testCase[int16]{
		{"0", 0},
		{"32767", 32767},
		{"-32768", -32768},
		{"32768", 0},  // Overflow
		{"-32769", 0}, // Underflow
		{"abc", 0},    // Invalid input
	}

	for _, tc := range testCasesInt16 {
		result := AtoI[int16](tc.input)
		if result != tc.expected {
			t.Errorf("AtoI[int16](%q) = %d; want %d", tc.input, result, tc.expected)
		}
	}

	// Test cases for uint16
	testCasesUint16 := []testCase[uint16]{
		{"0", 0},
		{"65535", 65535},
		{"65536", 0}, // Overflow
		{"-1", 0},    // Negative number for unsigned
		{"abc", 0},   // Invalid input
	}

	for _, tc := range testCasesUint16 {
		result := AtoI[uint16](tc.input)
		if result != tc.expected {
			t.Errorf("AtoI[uint16](%q) = %d; want %d", tc.input, result, tc.expected)
		}
	}

	// Test cases for int32
	testCasesInt32 := []testCase[int32]{
		{"0", 0},
		{"2147483647", 2147483647},
		{"-2147483648", -2147483648},
		{"2147483648", 0},  // Overflow
		{"-2147483649", 0}, // Underflow
		{"abc", 0},         // Invalid input
	}

	for _, tc := range testCasesInt32 {
		result := AtoI[int32](tc.input)
		if result != tc.expected {
			t.Errorf("AtoI[int32](%q) = %d; want %d", tc.input, result, tc.expected)
		}
	}

	// Test cases for uint32
	testCasesUint32 := []testCase[uint32]{
		{"0", 0},
		{"4294967295", 4294967295},
		{"4294967296", 0}, // Overflow
		{"-1", 0},         // Negative number for unsigned
		{"abc", 0},        // Invalid input
	}

	for _, tc := range testCasesUint32 {
		result := AtoI[uint32](tc.input)
		if result != tc.expected {
			t.Errorf("AtoI[uint32](%q) = %d; want %d", tc.input, result, tc.expected)
		}
	}

	// Test cases for int64
	testCasesInt64 := []testCase[int64]{
		{"0", 0},
		{"9223372036854775807", 9223372036854775807},
		{"-9223372036854775808", -9223372036854775808},
		{"9223372036854775808", 0},  // Overflow
		{"-9223372036854775809", 0}, // Underflow
		{"abc", 0},                  // Invalid input
	}

	for _, tc := range testCasesInt64 {
		result := AtoI[int64](tc.input)
		if result != tc.expected {
			t.Errorf("AtoI[int64](%q) = %d; want %d", tc.input, result, tc.expected)
		}
	}

	// Test cases for uint64
	testCasesUint64 := []testCase[uint64]{
		{"0", 0},
		{"18446744073709551615", 18446744073709551615},
		{"18446744073709551616", 0}, // Overflow
		{"-1", 0},                   // Negative number for unsigned
		{"abc", 0},                  // Invalid input
	}

	for _, tc := range testCasesUint64 {
		result := AtoI[uint64](tc.input)
		if result != tc.expected {
			t.Errorf("AtoI[uint64](%q) = %d; want %d", tc.input, result, tc.expected)
		}
	}

	// Test cases for int (architecture-dependent)
	testCasesInt := []testCase[int]{
		{"0", 0},
		{strconv.FormatInt(int64(math.MaxInt), 10), int(math.MaxInt)},
		{strconv.FormatInt(int64(math.MinInt), 10), int(math.MinInt)},
		{strconv.FormatUint(uint64(math.MaxInt)+1, 10), 0}, // Overflow
		{"abc", 0}, // Invalid input
	}

	for _, tc := range testCasesInt {
		result := AtoI[int](tc.input)
		if result != tc.expected {
			t.Errorf("AtoI[int](%q) = %d; want %d", tc.input, result, tc.expected)
		}
	}

	// Test cases for uint (architecture-dependent)
	testCasesUint := []testCase[uint]{
		{"0", 0},
		{strconv.FormatUint(uint64(math.MaxUint), 10), uint(math.MaxUint)},
		{strconv.FormatUint(uint64(math.MaxUint), 10) + "0", 0}, // Overflow
		{"-1", 0},  // Negative number for unsigned
		{"abc", 0}, // Invalid input
	}

	for _, tc := range testCasesUint {
		result := AtoI[uint](tc.input)
		if result != tc.expected {
			t.Errorf("AtoI[uint](%q) = %d; want %d", tc.input, result, tc.expected)
		}
	}
}

func ExampleAtoI() {
	i := AtoI[uint8]("23")
	fmt.Println(i)

	// Output: 23
}
