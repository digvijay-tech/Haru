// Helper function
package interpreter

import (
	"regexp"
	"strings"
)

func isInt(s string) bool {
	match, _ := regexp.MatchString(`^-?[0-9]+$`, s)
	return match
}

func isFloat(s string) bool {
	match, _ := regexp.MatchString(`^-?[0-9]+\.[0-9]+$`, s)
	return match
}

func isByte(s string) bool {
	return strings.HasPrefix(s, "0b")
}

func isString(s string) bool {
	return strings.HasPrefix(s, "\"") || strings.HasPrefix(s, "'")
}

func stripQuotes(s string) string {
	if len(s) >= 2 && (s[0] == '"' || s[0] == '\'') {
		return s[1 : len(s)-1]
	}
	return s
}
