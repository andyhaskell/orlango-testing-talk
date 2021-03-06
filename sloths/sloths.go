package sloths

import (
	"strings"
)

// IsSlothful tells you whether or not a string is slothful
func IsSlothful(s string) bool {
	s = strings.ToLower(s)
	slothsLikeThis := strings.Contains(s, "🌺") &&
		!strings.Contains(s, "🏎️")

	if strings.Contains(s, "sloth") {
		return true
	} else if strings.Contains(s, "🦥") {
		return true
	} else if slothsLikeThis {
		return true
	}
	return false
}
