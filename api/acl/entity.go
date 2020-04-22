package entity

// Entity acl table entity
type Entity struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Lv          int    `json:"lv"`
	SortNo      int    `json:"sortNo"`
	Enable      bool   `json:"enable"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}

// Col get acl table column name
func (entity *Entity) Col() Col {
	return Col{
		ID:          "id",
		Code:        "code",
		Name:        "name",
		Type:        "type",
		Lv:          "lv",
		SortNo:      "sort_no",
		Enable:      "enable",
		CreatedDate: "created_date",
		UpdatedDate: "updated_date",
	}
}

// Col role_acl table column name
type Col struct {
	ID          string
	Code        string
	Name        string
	Type        string
	Lv          string
	SortNo      string
	Enable      string
	CreatedDate string
	UpdatedDate string
}

// Table acl table name
func (entity *Entity) Table() string {
	return "acl"
}

// PK acl table pk column name
func (entity *Entity) PK() string {
	return "id"
}
