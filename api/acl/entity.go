package entity

// Table acl table name
const Table string = "acl"

// Entity acl table entity
type Entity struct {
	ID          int    `json:"id" db:"id"`
	Code        string `json:"code" db:"code"`
	Name        string `json:"name" db:"name"`
	Type        string `json:"type" db:"type"`
	Lv          int    `json:"lv" db:"lv"`
	SortNo      int    `json:"sortNo" db:"sort_no"`
	Enable      bool   `json:"enable" db:"enable"`
	CreatedDate string `json:"creatDate" db:"created_date"`
	UpdatedDate string `json:"updateDate" db:"updated_date"`
}

// Table acl table name
func (entity *Entity) Table() string {
	return "acl"
}

// PK acl table pk column name
func (entity *Entity) PK() string {
	return "id"
}
