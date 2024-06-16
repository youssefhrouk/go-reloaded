
package reloaded

import (
	"strings"
	"unicode"
)

func IsAlpha(s string) bool {
	for _, char := range s {
		if !unicode.IsLetter(char) {
			return false
		}
	}
	return true
}
func Capitalize(s string) string {
	s = strings.ToLower(s)
	for i, v := range s {
		if i == 0 {
			s = strings.ToUpper(string(v)) + s[i+1:]
		} else {
			if !IsAlpha(string(v)) && !IsAlpha(string(s[i-1])) {
				if i != len(s)-1 {
					s = s[:i] + strings.ToUpper(string(v)) + s[i+1:]
				} else {
					s = s[:i] + strings.ToUpper(string(v))
				}
			}
		}
	}
	return s
}
