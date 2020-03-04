package datahandler

import (
	"reflect"
	"unicode"
)

// GetColName get entity column name
func GetColName(entity interface{}, key string) string {
	s, ok := reflect.TypeOf(entity).Elem().FieldByName(key)
	if ok {
		col := s.Tag.Get("db")
		if len(col) > 0 {
			return col
		}
	}

	return Camel2Snake(key)
}

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
		return result[1:]
	}
	return result
}
