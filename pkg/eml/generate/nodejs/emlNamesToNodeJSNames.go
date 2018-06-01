package nodejs

import (
	"strings"
	"unicode"
)

// ToNodeJsClassName ...
func ToNodeJsClassName(s string) string {
	return strings.Replace(s, " ", "", -1)
}

// FirstCharToLower ...
func FirstCharToLower(s string) string {
	if len(s) == 0 {
		return s
	}
	a := []rune(s)
	a[0] = unicode.ToLower(a[0])
	return (string(a))
}
