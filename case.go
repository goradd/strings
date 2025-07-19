package strings

import (
	"strings"
	"unicode"
)

// KebabToCamel convert string s from kebab-case to CamelCase. The initial case of s does not affect the output.
func KebabToCamel(s string) string {
	// This depends on Map walking the string from start to end, which
	// really needs to happen anyway in unicode strings.
	capitalize := true
	return strings.Map(func(r rune) rune {
		if r == '-' {
			capitalize = true
			return -1
		} else if capitalize {
			capitalize = false
			return unicode.ToTitle(r)
		} else {
			return unicode.ToLower(r)
		}
	}, s)
}

// SnakeToKebab converts s from snake_case to kebab-case. Only the underscores are impacted, the case of the rest of s
// is passed through unchanged.
func SnakeToKebab(s string) string {
	return strings.Replace(s, "_", "-", -1)
}

// SnakeToCamel converts s from snake_case to CamelCase.
func SnakeToCamel(s string) string {
	k := SnakeToKebab(s)
	return KebabToCamel(k)
}

// CamelToKebab converts capitalize from CamelCase to kebab-case.
// If it encounters a character that is not legitimate camel case,
// it ignores it (like numbers, spaces, etc.).
// Runs of upper case letters are treated as one word.
func CamelToKebab(camelCase string) string {
	return camelToKebabOrSnake(camelCase, '-')
}

// CamelToSnake converts camelCase from CamelCase to snake_case.
// If it encounters a character that is not legitimate camel case,
// it ignores it (like numbers, spaces, etc.).
// Runs of upper case letters are treated as one word.
// A run of upper case, followed by lower case letters will be treated
// as if the final character in the upper case run belongs with the lower case
// letters.
func CamelToSnake(camelCase string) string {
	return camelToKebabOrSnake(camelCase, '_')
}

func camelToKebabOrSnake(camelCase string, replacement rune) string {
	var kebabCase []rune
	var inUpper bool

	for i, r := range camelCase {
		if unicode.IsLetter(r) {
			if unicode.IsUpper(r) {
				if i > 0 && !inUpper {
					kebabCase = append(kebabCase, replacement)
				}
				kebabCase = append(kebabCase, unicode.ToLower(r))
				inUpper = true

			} else {
				if inUpper {
					// switching from upper to lower, if we were in an upper run
					// we need to add a hyphen in front of the last rune
					if len(kebabCase) > 1 && kebabCase[len(kebabCase)-2] != replacement {
						r2 := kebabCase[len(kebabCase)-1]
						kebabCase[len(kebabCase)-1] = replacement
						kebabCase = append(kebabCase, r2)
					}
				}
				kebabCase = append(kebabCase, r)
				inUpper = false
			}
		}
	}

	return string(kebabCase)
}

// Decap returns a new string with the first character in the string set to its lower case equivalent,
// and subsequent characters that are capitalized also set to lower case until it encounters a lower case letter.
func Decap(s string) string {
	if s == "" {
		return ""
	}
	var b strings.Builder
	var foundLower bool
	for i, r := range s {
		l := unicode.ToLower(r)
		if i == 0 {
			b.WriteRune(l)
		} else if l != r && !foundLower {
			b.WriteRune(l)
		} else {
			foundLower = true
			b.WriteRune(r)
		}
	}
	return b.String()
}

// Title is a more advanced titling operation. It will convert underscores to spaces, and add spaces to CamelCase
// words.
func Title(s string) string {
	s = strings.TrimSpace(strings.Title(strings.Replace(s, "_", " ", -1)))
	if len(s) <= 1 {
		return s
	}

	newString := s[0:1]
	l := strings.ToLower(s)
	i := 1
	for i < len(s) {
		if l[i] != s[i] && s[i-1:i] != " " {
			// is a capital.
			newString += " "
			// Group capitals until we get a lower case
			for i < len(s)-1 && l[i] != s[i] && l[i+1] != s[i+1] {
				newString += s[i : i+1]
				i++
			}
		}
		newString += s[i : i+1]
		i++
	}
	return newString
}

// Camel converts a string to camel case.
func Camel(s string) string {
	s = strings.TrimSpace(strings.Title(strings.Replace(s, "_", " ", -1)))
	if len(s) <= 1 {
		return s
	}

	newString := s[0:1]
	l := strings.ToLower(s)
	for i := 1; i < len(s); i++ {
		if l[i] != s[i] && s[i-1:i] != " " {
			newString += " "
		}
		newString += s[i : i+1]
	}
	return newString
}

// EqualCaseInsensitive is a synonym for the strings package EqualFold which provides unicode compliant case-insensitive
// comparison.
func EqualCaseInsensitive(s1, s2 string) bool {
	return strings.EqualFold(s1, s2)
}

func IsSnake(s string) bool {
	for _, r := range s {
		if !((unicode.IsLetter(r) && unicode.IsLower(r)) || r == '_' || unicode.IsNumber(r)) {
			return false
		}
	}
	return true
}
