package strings

import (
	"testing"
)

// TestIsASCII tests the IsASCII function.
func TestIsASCII(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Hello, World!", true},
		{"こんにちは", false},   // Japanese characters
		{"12345", true},    // Numeric ASCII characters
		{"", true},         // Empty string
		{"abc\x80", false}, // Non-ASCII byte in string
		{"\x7F\x7E", true}, // Edge ASCII values
	}

	for _, tt := range tests {
		result := IsASCII(tt.input)
		if result != tt.expected {
			t.Errorf("IsASCII(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

// TestIsUTF8 tests the IsUTF8 function.
func TestIsUTF8(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Hello, World!", true},
		{"こんにちは", true},         // Valid UTF-8 Japanese characters
		{"\xff\xfe\xfd", false}, // Invalid UTF-8 byte sequence
		{"12345", true},         // ASCII is valid UTF-8
		{"", true},              // Empty string
	}

	for _, tt := range tests {
		result := IsUTF8(tt.input)
		if result != tt.expected {
			t.Errorf("IsUTF8(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

// TestIsUTF8Bytes tests the IsUTF8Bytes function.
func TestIsUTF8Bytes(t *testing.T) {
	tests := []struct {
		input    []byte
		expected bool
	}{
		{[]byte("Hello, World!"), true},
		{[]byte("こんにちは"), true},           // Valid UTF-8 Japanese characters
		{[]byte{0xff, 0xfe, 0xfd}, false}, // Invalid UTF-8 byte sequence
		{[]byte("12345"), true},           // ASCII is valid UTF-8
		{[]byte{}, true},                  // Empty byte slice
	}

	for _, tt := range tests {
		result := IsUTF8Bytes(tt.input)
		if result != tt.expected {
			t.Errorf("IsUTF8Bytes(%v) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

// TestIsInt tests the IsInt function.
func TestIsInt(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123", true},
		{"-123", true},
		{"+123", true},
		{"123.45", false}, // Not an integer
		{"abc", false},    // Not a number
		{"", false},       // Empty string
	}

	for _, tt := range tests {
		result := IsInt(tt.input)
		if result != tt.expected {
			t.Errorf("IsInt(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

// TestIsFloat tests the IsFloat function.
func TestIsFloat(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"123.45", true},
		{"-123.45", true},
		{"+123.45", true},
		{"123", true},  // Integers are valid floats
		{"abc", false}, // Not a number
		{"", false},    // Empty string
	}

	for _, tt := range tests {
		result := IsFloat(tt.input)
		if result != tt.expected {
			t.Errorf("IsFloat(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}

// TestStripNewlines tests the StripNewlines function.
func TestStripNewlines(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello\nWorld", "HelloWorld"},
		{"Hello\r\nWorld", "HelloWorld"},
		{"Hello World", "Hello World"}, // No newlines
		{"\n\r\n\r", ""},               // Only newlines
		{"", ""},                       // Empty string
	}

	for _, tt := range tests {
		result := StripNewlines(tt.input)
		if result != tt.expected {
			t.Errorf("StripNewlines(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

// TestStripNulls tests the StripNulls function.
func TestStripNulls(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"Hello\x00World", "HelloWorld"},
		{"Hello World", "Hello World"}, // No null characters
		{"\x00\x00\x00", ""},           // Only null characters
		{"", ""},                       // Empty string
	}

	for _, tt := range tests {
		result := StripNulls(tt.input)
		if result != tt.expected {
			t.Errorf("StripNulls(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}

// TestHasNull tests the HasNull function.
func TestHasNull(t *testing.T) {
	tests := []struct {
		input    string
		expected bool
	}{
		{"Hello\x00World", true},
		{"Hello World", false}, // No null characters
		{"\x00", true},         // Single null character
		{"", false},            // Empty string
	}

	for _, tt := range tests {
		result := HasNull(tt.input)
		if result != tt.expected {
			t.Errorf("HasNull(%q) = %v; want %v", tt.input, result, tt.expected)
		}
	}
}
