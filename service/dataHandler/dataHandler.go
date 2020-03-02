package dataHandler

import (
	"unicode"
)

// Camel2Snake convert string from Camel-Case to Snake-Case
func Camel2Snake(key string) string {
	if len(key) < 1 {
		return ""
	}

	result := ""
	for _, r := range key {
		if unicode.IsUpper(r) {
			result += "_" + string(unicode.ToLower(r))
		} else {
			result += string(r)
		}
	}

	if result[0:1] == "_" {
		return result[0:1]
	}
	return result
}
