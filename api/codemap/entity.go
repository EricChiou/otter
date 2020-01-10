package codemap

// Table codemap table name
const Table string = "codemap"

// PK codemap table pk column name
const PK string = "id"

// Entity role table entity
type Entity struct {
	ID          int    `json:"id"`
	Type        string `json:"type"`
	Code        string `json:"code"`
	Name        string `json:"name"`
	SortNo      int    `json:"sortNo"`
	Enable      bool   `json:"enable"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}

// Col role table column name
var Col col = col{
	ID:          "id",
	Type:        "type",
	Code:        "code",
	Name:        "name",
	SortNo:      "sort_no",
	Enable:      "enable",
	CreatedDate: "created_date",
	UpdatedDate: "updated_date",
}

type col struct {
	ID          string
	Type        string
	Code        string
	Name        string
	SortNo      string
	Enable      string
	CreatedDate string
	UpdatedDate string
}
