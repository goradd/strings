package strings

import (
	"fmt"
	"testing"
)

func ExampleKebabToCamel() {
	a := KebabToCamel("abc-def")
	fmt.Println(a)
	//Output: AbcDef
}

func TestKebabToCamel(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"a", "a", "A"},
		{"-a", "-a", "A"},
		{"b-a", "b-a", "BA"},
		{"123-abc", "123-abc", "123Abc"},
		{"this-that", "this-that", "ThisThat"},
		{"This-THAT", "This-THAT", "ThisThat"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := KebabToCamel(tt.s); got != tt.want {
				t.Errorf("KebabToCamel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleSnakeToKebab() {
	a := SnakeToKebab("abc_def")
	fmt.Println(a)
	//Output: abc-def
}

func TestSnakeToKebab(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"a_b", "a_b", "a-b"},
		{"-b", "-b", "-b"},
		{"_b", "_b", "-b"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SnakeToKebab(tt.s); got != tt.want {
				t.Errorf("SnakeToKebab() = %v, want %v", got, tt.want)
			}
		})
	}
}

func ExampleCamelToKebab() {
	a := CamelToKebab("AbcDef")
	fmt.Println(a)
	b := CamelToKebab("AbcDEFghi")
	fmt.Println(b)
	//Output: abc-def
	//abc-de-fghi
}

func ExampleCamelToSnake() {
	a := CamelToSnake("AbcDef")
	fmt.Println(a)
	b := CamelToSnake("AbcDEFghi")
	fmt.Println(b)
	//Output: abc_def
	//abc_de_fghi
}

func TestCamelToKebab(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"a", "a", "a"},
		{"A", "A", "a"},
		{"ab", "ab", "ab"},
		{"AB", "AB", "ab"},
		{"Ab", "Ab", "ab"},
		{"aB", "aB", "a-b"},
		{"Abc", "Abc", "abc"},
		{"AbC", "AbC", "ab-c"},
		{"ABc", "ABc", "a-bc"},
		{"a1b", "a1b", "ab"},
		{"ABC", "ABC", "abc"},
		{"ABCd", "ABCd", "ab-cd"},
		{"AbCdE", "ABCdE", "ab-cd-e"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CamelToKebab(tt.s); got != tt.want {
				t.Errorf("SnakeToKebab() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDecap(t *testing.T) {
	if Decap("") != "" {
		t.Fail()
	}
	if Decap("A") != "a" {
		t.Fail()
	}
	if Decap("AbcDef") != "abcDef" {
		t.Fail()
	}
	if Decap("Value") != "id" {
		t.Fail()
	}
	if Decap("IDs") != "ids" {
		t.Fail()
	}
	if Decap("IDsFor") != "idsFor" {
		t.Fail()
	}

}

func ExampleTitle() {
	a := Title("do_i_seeYou")
	fmt.Println(a)
	//Output: Do I See You
}

func TestTitle(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want string
	}{
		{"empty", "", ""},
		{"i", "i", "I"},
		{"iJ", "iJ", "I J"},
		{"i_j", "i_j", "I J"},
		{"iJK", "iJK", "I JK"},
		{"i_J_k", "i_J_k", "I J K"},
		{"ManagerID", "ManagerID", "Manager ID"},
		{"BobTheGrocer", "BobTheGrocer", "Bob The Grocer"},
		{"ILike Kiwis", "ILike Kiwis", "I Like Kiwis"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Title(tt.s); got != tt.want {
				t.Errorf("Title() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestEqualCaseInsensitive(t *testing.T) {
	tests := []struct {
		name string
		s1   string
		s2   string
		want bool
	}{
		{"match", "aBc", "Abc", true},
		{"nomatch", "aBd", "Abc", false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := EqualCaseInsensitive(tt.s1, tt.s2); got != tt.want {
				t.Errorf("EqualCaseInsensitive() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsSnake(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected bool
	}{
		{
			name:     "Valid snake_case string",
			input:    "valid_snake_case",
			expected: true,
		},
		{
			name:     "Contains uppercase letters",
			input:    "Invalid_Snake_Case",
			expected: false,
		},
		{
			name:     "Contains special characters",
			input:    "snake_case@123",
			expected: false,
		},
		{
			name:     "Contains spaces",
			input:    "snake case",
			expected: false,
		},
		{
			name:     "Contains only numbers",
			input:    "12345",
			expected: true,
		},
		{
			name:     "Empty string",
			input:    "",
			expected: true,
		},
		{
			name:     "Contains consecutive underscores",
			input:    "snake__case",
			expected: true,
		},
		{
			name:     "Valid snake_case with numbers",
			input:    "snake_case_123",
			expected: true,
		},
		{
			name:     "Valid with trailing underscore",
			input:    "snake_case_",
			expected: true,
		},
		{
			name:     "Single underscore",
			input:    "_",
			expected: true,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := IsSnake(tt.input)
			if result != tt.expected {
				t.Errorf("IsSnake(%q) = %v; want %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestSnakeToCamel(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"", ""},
		{"hello_world", "HelloWorld"},
		{"a_b_c", "ABC"},
		{"this_is_a_test", "ThisIsATest"},
		{"snake_case", "SnakeCase"},
		{"alreadycamel", "Alreadycamel"}, // Will depend on behavior
		{"_leading_underscore", "LeadingUnderscore"},
		{"trailing_underscore_", "TrailingUnderscore"},
		{"__multiple__underscores__", "MultipleUnderscores"},
		{"with_numbers_123", "WithNumbers123"},
		{"a", "A"},
	}

	for _, tt := range tests {
		result := SnakeToCamel(tt.input)
		if result != tt.expected {
			t.Errorf("SnakeToCamel(%q) = %q; want %q", tt.input, result, tt.expected)
		}
	}
}
