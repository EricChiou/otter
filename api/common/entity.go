package common

import (
	"reflect"

	"otter/service/dataHandler"
)

// BaseEntity base entity function
type BaseEntity struct{}

// Col get entity column name
func (entity *BaseEntity) Col(key string) string {
	s, ok := reflect.TypeOf(entity).Elem().FieldByName(key)
	if ok {
		col := s.Tag.Get("db")
		if len(col) > 0 {
			return col
		}
	}

	return dataHandler.Camel2Snake(key)
}
