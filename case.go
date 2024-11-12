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