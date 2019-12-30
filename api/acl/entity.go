package entity

// Table acl table name
const Table string = "acl"

// Col acl table column name
var Col col = col{
	PK:          "code",
	Code:        "code",
	Name:        "name",
	Type:        "type",
	Lv:          "lv",
	SortNo:      "sort_no",
	Enable:      "enable",
	CreatedDate: "created_date",
	UpdatedDate: "updated_date",
}

// Entity acl table entity
type Entity struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Lv          int    `json:"lv"`
	SortNo      int    `json:"sortNo"`
	Enable      bool   `json:"enable"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}

type col struct {
	PK          string
	Code        string
	Name        string
	Type        string
	Lv          string
	SortNo      string
	Enable      string
	CreatedDate string
	UpdatedDate string
}
