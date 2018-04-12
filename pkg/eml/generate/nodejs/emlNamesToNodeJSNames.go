package nodejs

import "strings"

// ToNodeJsClassName ...
func ToNodeJsClassName(s string) string {
	return strings.Replace(s, " ", "", -1)
}
