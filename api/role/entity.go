package role

// Table role table name
const Table string = "role"

// PK role table pk column name
const PK string = "id"

// Entity role table entity
type Entity struct {
	ID          int    `json:"id"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	Lv          int    `json:"lv"`
	SortNo      int    `json:"sortNo"`
	Enable      bool   `json:"enable"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}

// Col role table column name
var Col col = col{
	ID:          "id",
	Code:        "code",
	Name:        "name",
	Lv:          "lv",
	SortNo:      "sort_no",
	Enable:      "enable",
	CreatedDate: "created_date",
	UpdatedDate: "updated_date",
}

type col struct {
	ID          string
	Code        string
	Name        string
	Lv          string
	SortNo      string
	Enable      string
	CreatedDate string
	UpdatedDate string
}
