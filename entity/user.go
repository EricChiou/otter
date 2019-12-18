package entity

import (
	"reflect"
)

// User table struct
type User struct {
	ID          int    `json:"id" db:"id"`
	Email       string `json:"email" db:"email"`
	Pwd         string `json:"pwd" db:"pwd"`
	Name        string `json:"name" db:"name"`
	Identity    string `json:"identity" db:"identity"`
	Active      bool   `json:"active" db:"active"`
	CreatedDate string `json:"creatDate" db:"created_date"`
	UpdatedDate string `json:"updateDate" db:"updated_date"`
}

// Col get column name by input a struct key
func (user *User) Col(key string) string {
	types := reflect.TypeOf(user)
	for i := 0; i < types.NumField(); i++ {
		if types.Field(i).Name == key {
			return types.Field(i).Tag.Get("db")
		}
	}

	return ""
}
