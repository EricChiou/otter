package user

import (
	"otter/api/common"
)

// Entity user table entity
type Entity struct {
	common.BaseEntity
	ID          int    `json:"id" db:"id"`
	Email       string `json:"email" db:"email"`
	Pwd         string `json:"pwd" db:"pwd"`
	Name        string `json:"name" db:"name"`
	Role        string `json:"role" db:"role"`
	Active      bool   `json:"active" db:"active"`
	CreatedDate string `json:"creatDate" db:"created_date"`
	UpdatedDate string `json:"updateDate" db:"updated_date"`
}

// Table user table name
func (entity *Entity) Table() string {
	return "user"
}

// PK user table pk column name
func (entity *Entity) PK() string {
	return "id"
}
