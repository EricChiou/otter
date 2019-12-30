package entity

// Table role table name
const Table string = "role"

// Col role table column name
var Col col = col{
	PK:          "code",
	Code:        "code",
	Name:        "name",
	Lv:          "lv",
	SortNo:      "sort_no",
	Enable:      "enable",
	CreatedDate: "created_date",
	UpdatedDate: "updated_date",
}

// Entity role table entity
type Entity struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
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
	Lv          string
	SortNo      string
	Enable      string
	CreatedDate string
	UpdatedDate string
}
