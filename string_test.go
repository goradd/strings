package strings

import (
	"fmt"
	"testing"
)

func TestIndent(t *testing.T) {
	if Indent("a\nb\nc") != "\ta\n\tb\n\tc" {
		t.Fail()
	}
	if Indent("\na\nb\nc") != "\t\n\ta\n\tb\n\tc" {
		t.Fail()
	}
	if Indent("a\nb\nc\n") != "\ta\n\tb\n\tc\n" {
		t.Fail()
	}
}

func TestHasOnlyLetters(t *testing.T) {
	if HasOnlyLetters("a-b") {
		t.Fail()
	}
	if !HasOnlyLetters("abc") {
		t.Fail()
	}
	if HasOnlyLetters("123") {
		t.Fail()
	}
}

func TestStartsWith(t *testing.T) {
	type args struct {
	}
	tests := []struct {
		name      string
		s         string
		beginning string
		want      bool
	}{
		{"same with dot", ".45", ".45", true},
		{"short", "a", "a", true},
		{"short2", "abc", "a", true},
		{"short3", "234f asd fa", "a", false},
		{"mid", "234f abc fa", "abc", false},
		{"smaller", "ab", "abc", false},
		{"smaller2", "abc", "ab", true},
		{"none", "", "abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := StartsWith(tt.s, tt.beginning); got != tt.want {
				t.Errorf("StartsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEndsWith(t *testing.T) {
	tests := []struct {
		name   string
		s      string
		ending string
		want   bool
	}{
		{"same", ".45", ".45", true},
		{"a", "a", "a", true},
		{"long", "234f asd fa", "a", true},
		{"long2", "asdfsaf sdabc", "abc", true},
		{"too short", "bc", "abc", false},
		{"empty", "", "abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EndsWith(tt.s, tt.ending); got != tt.want {
				t.Errorf("EndsWith() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleConnect() {
	a := Connect("+", "this", "", "that")
	fmt.Println(a)
	//Output: this+that
}

func TestConnect(t *testing.T) {
	type args struct {
		sep   string
		items []string
	}
	tests := []struct {
		name  string
		sep   string
		items []string
		want  string
	}{
		{"empty", "", []string{""}, ""},
		{"1", "+", []string{"this"}, "this"},
		{"2", "+", []string{"this", "that"}, "this+that"},
		{"empty sep", "", []string{"this", "that"}, "thisthat"},
		{"empty item", "+", []string{"this", "", "that"}, "this+that"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Connect(tt.sep, tt.items...); got != tt.want {
				t.Errorf("JoinContent() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIf(t *testing.T) {
	if If(true, "a", "b") != "a" {
		t.Fail()
	}
	if If(false, "a", "b") != "b" {
		t.Fail()
	}
}

func TestContainsAnyStrings(t *testing.T) {
	tests := []struct {
		name     string
		haystack string
		needles  []string
		want     bool
	}{
		{"empty", "", []string{}, false},
		{"empty2", "", []string{"a", "b"}, false},
		{"a", "a", []string{"a", "b"}, true},
		{"b", "b", []string{"a", "b"}, true},
		{"abc", "abc", []string{"h", "bc"}, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ContainsAnyStrings(tt.haystack, tt.needles...); got != tt.want {
				t.Errorf("ContainsAnyStrings() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestHasCharType(t *testing.T) {
	type args struct {
		s          string
		wantUpper  bool
		wantLower  bool
		wantDigit  bool
		wantPunc   bool
		wantSymbol bool
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{"lower", args{"a", false, true, false, false, false}, true},
		{"lowerFail", args{"A", false, true, false, false, false}, false},
		{"upper", args{"A", true, false, false, false, false}, true},
		{"upperFail", args{"a", true, false, true, false, false}, false},
		{"digit", args{"1", false, false, true, false, false}, true},
		{"digitFail", args{"A", false, false, true, false, false}, false},
		{"punc", args{",", false, false, false, true, false}, true},
		{"puncFail", args{"a", false, false, false, true, false}, false},
		{"symbol", args{"$", false, false, false, false, true}, true},
		{"symbolFail", args{",", false, false, false, false, true}, false},
		{"mult1", args{"aA", true, true, false, false, false}, true},
		{"mult1Fail", args{"a1", true, true, false, false, false}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := HasCharType(tt.args.s, tt.args.wantUpper, tt.args.wantLower, tt.args.wantDigit, tt.args.wantPunc, tt.args.wantSymbol); got != tt.want {
				t.Errorf("HasCharType() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestReplaceStrings tests the ReplaceStrings function for various cases.
func TestReplaceStrings(t *testing.T) {
	tests := []struct {
		input       string
		searchList  []string
		replaceList []string
		expected    string
	}{
		// Basic replacement
		{"hello world", []string{"hello", "world"}, []string{"hi", "planet"}, "hi planet"},
		// No replacements
		{"hello world", []string{"foo", "bar"}, []string{"baz", "qux"}, "hello world"},
		// Partial replacement
		{"hello world", []string{"world"}, []string{"planet"}, "hello planet"},
		// Empty input
		{"", []string{"hello"}, []string{"hi"}, ""},
		// Empty search and replace lists
		{"hello world", []string{}, []string{}, "hello world"},
		// Overlapping search strings
		{"abcabcabc", []string{"abc", "a"}, []string{"xyz", "x"}, "xyzxyzxyz"},
		// Search list longer than replace list
		{"test abc", []string{"test", "abc", "xyz"}, []string{"TEST", "ABC", ""}, "TEST ABC"},
		// Replace list longer than search list (ignored extra)
		{"test abc", []string{"test", "abc"}, []string{"TEST", "ABC", "EXTRA"}, "TEST ABC"},
	}

	for _, tt := range tests {
		result := ReplaceStrings(tt.input, tt.searchList, tt.replaceList)
		if result != tt.expected {
			t.Errorf("ReplaceStrings(%q, %v, %v) = %q; want %q", tt.input, tt.searchList, tt.replaceList, result, tt.expected)
		}
	}
}

func TestPlural(t *testing.T) {
	tests := []struct {
		name  string
		input string
		want  string
	}{
		{"empty", "", ""},
		{"dog", "dog", "dogs"},
		{"person", "person", "people"},
		{"people", "people", "people"},
		{"group", "group", "groups"},
		{"big octopus", "big octopus", "big octopuses"},
		{"sheep", "sheep", "sheep"},
		{"bob person", "bob person", "bob people"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Plural(tt.input); got != tt.want {
				t.Errorf("Plural() = %v, want %v", got, tt.want)
			}
		})
	}
}

// TestBetween tests the Between function
func TestBetween(t *testing.T) {
	tests := []struct {
		name     string
		s        string
		left     string
		right    string
		expected string
	}{
		{
			name:     "Valid substring between markers",
			s:        "Hello [world]!",
			left:     "[",
			right:    "]",
			expected: "world",
		},
		{
			name:     "No left marker",
			s:        "Hello world!",
			left:     "[",
			right:    "]",
			expected: "Hello world!",
		},
		{
			name:     "No right marker",
			s:        "Hello [world!",
			left:     "[",
			right:    "]",
			expected: "Hello [world!",
		},
		{
			name:     "No markers",
			s:        "Hello world!",
			left:     "{",
			right:    "}",
			expected: "Hello world!",
		},
		{
			name:     "Empty input string",
			s:        "",
			left:     "[",
			right:    "]",
			expected: "",
		},
		{
			name:     "Empty left marker",
			s:        "Hello [world]!",
			left:     "",
			right:    "]",
			expected: "Hello [world]!",
		},
		{
			name:     "Empty right marker",
			s:        "Hello [world]!",
			left:     "[",
			right:    "",
			expected: "Hello [world]!",
		},
		{
			name:     "Markers overlap",
			s:        "Hello [[[world]]]",
			left:     "[[",
			right:    "]]",
			expected: "[world]",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Between(tt.s, tt.left, tt.right)
			if result != tt.expected {
				t.Errorf("Between(%q, %q, %q) = %q; want %q", tt.s, tt.left, tt.right, result, tt.expected)
			}
		})
	}
}
