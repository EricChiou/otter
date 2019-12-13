package checkparam

import (
	"fmt"
	"reflect"
)

// Check check request parameters
func Check(params ...interface{}) bool {
	for _, param := range params {
		name := reflect.TypeOf(param).Name()
		if name == "string" {
			str := fmt.Sprintf("%v", param)
			if len(str) == 0 {
				return false
			}
		} else if name == "int" || name == "int8" || name == "int16" || name == "int32" || name == "int64" ||
			name == "uint" || name == "uint8" || name == "uint16" || name == "uint32" || name == "uint64" || name == "uintptr" ||
			name == "float32" || name == "float64" ||
			name == "complex64" || name == "complex128" {
			if param == 0 {
				return false
			}
		} else if name == "bool" {
			return true
		} else {
			if param == nil {
				return false
			}
		}
	}
	return true
}
