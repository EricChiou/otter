package entity

// Entity acl table entity
type Entity struct {
	ID          int    `json:"id,omitempty"`
	Code        string `json:"code,omitempty"`
	Name        string `json:"name,omitempty"`
	Type        string `json:"type,omitempty"`
	Lv          int    `json:"lv,omitempty"`
	SortNo      int    `json:"sortNo,omitempty"`
	Enable      bool   `json:"enable,omitempty"`
	CreatedDate string `json:"creatDate,omitempty"`
	UpdatedDate string `json:"updateDate,omitempty"`
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
