package strings

import (
	"github.com/gedex/inflector"
	"strings"
	"unicode"
)

// StartsWith returns true if the string begins with the beginning string.
func StartsWith(s string, beginning string) bool {
	return strings.HasPrefix(s, beginning)
}

// EndsWith returns true if the string ends with the ending string.
func EndsWith(s string, ending string) bool {
	return strings.HasSuffix(s, ending)
}

// Indent will indent every line of the string with a tab
func Indent(s string) string {
	s = "\t" + strings.Replace(s, "\n", "\n\t", -1)
	return strings.TrimRight(s, "\t")
}

// HasOnlyLetters will return false if any of the characters in the string do not pass the unicode.IsLetter test.
func HasOnlyLetters(s string) bool {
	for _, r := range s {
		if !unicode.IsLetter(r) {
			return false
		}
	}
	return true
}

// Connect joins strings together with the separator sep. Only strings that are not empty strings are joined.
func Connect(sep string, items ...string) string {
	var l []string
	for _, i := range items {
		if i != "" {
			l = append(l, i)
		}
	}
	return strings.Join(l, sep)
}

// If is like the ternary operator ?. It returns the first string on true, and the second on false.
func If(cond bool, trueVal, falseVal string) string {
	if cond {
		return trueVal
	} else {
		return falseVal
	}
}

// ContainsAnyStrings returns true if the haystack contains any of the needles
func ContainsAnyStrings(haystack string, needles ...string) bool {
	for _, h := range needles {
		if strings.Contains(haystack, h) {
			return true
		}
	}
	return false
}

// HasCharType returns true if the given string has at least one of the selected char types.
func HasCharType(s string, wantUpper, wantLower, wantDigit, wantPunc, wantSymbol bool) bool {
	var hasUpper, hasLower, hasDigit, hasPunc, hasSymbol bool

	for _, c := range s {
		if !hasUpper && wantUpper && unicode.IsUpper(c) {
			hasUpper = true
		} else if !hasLower && wantLower && unicode.IsLower(c) {
			hasLower = true
		} else if !hasDigit && wantDigit && unicode.IsDigit(c) {
			hasDigit = true
		} else if !hasPunc && wantPunc && unicode.IsPunct(c) {
			hasPunc = true
		} else if !hasSymbol && wantSymbol && unicode.IsSymbol(c) {
			hasSymbol = true
		}

		if (!wantUpper || hasUpper) &&
			(!wantLower || hasLower) &&
			(!wantDigit || hasDigit) &&
			(!wantPunc || hasPunc) &&
			(!wantSymbol || hasSymbol) {
			return true
		}
	}
	return false
}

// ReplaceStrings replaces every string in the searchList with the matching string in the replaceList.
// Will panic if searchList len does not match replaceList len, or anything else goes wrong in the replacement.
func ReplaceStrings(s string, searchList []string, replaceList []string) string {
	var oldnew []string
	for i, s := range searchList {
		oldnew = append(oldnew, s, replaceList[i])
	}
	return ReplaceOldNew(s, oldnew...)
}

// ReplaceOldNew replaces every string in oldNew with the string following it, such that each pair of strings
// forms an old/new pair.
func ReplaceOldNew(s string, searchList ...string) string {
	repl := strings.NewReplacer(searchList...)
	return repl.Replace(s)
}

// Plural returns the plural version of the given string.
// This relies on a third party library, which may or may not be accurate. The goal is to
// handle the most common cases.
func Plural(s string) string {
	words := strings.Fields(s) // splits by any whitespace
	if len(words) == 0 {
		return ""
	}
	words[len(words)-1] = inflector.Pluralize(words[len(words)-1])
	return strings.Join(words, " ")
}

// Between returns the string between the left and right values in s.
// If left or right are not in s, all of s is returned.
func Between(s, left, right string) string {
	// Find the first and last single quotes
	start := strings.Index(s, left)
	end := strings.LastIndex(s, right)

	// If no single quotes are found or they don't form a valid pair, return the entire string
	if left == "" || right == "" || start == -1 || end == -1 || start == end {
		return s
	}

	// Return the substring between the quotes
	return s[start+len(left) : end]
}
