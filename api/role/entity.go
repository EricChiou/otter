package role

import (
	"otter/service/datahandler"
)

// Entity role table entity
type Entity struct {
	ID          int    `json:"id" db:"id"`
	Code        string `json:"code" db:"code"`
	Name        string `json:"name" db:"name"`
	Lv          int    `json:"lv" db:"lv"`
	SortNo      int    `json:"sortNo" db:"sort_no"`
	Enable      bool   `json:"enable" db:"enable"`
	CreatedDate string `json:"creatDate" db:"created_date"`
	UpdatedDate string `json:"updateDate" db:"updated_date"`
}

// Table role table name
func (entity *Entity) Table() string {
	return "role"
}

// PK role table pk column name
func (entity *Entity) PK() string {
	return "id"
}

// Col get entity column name
func (entity *Entity) Col(key string) string {
	return datahandler.GetColName(entity, key)
}
