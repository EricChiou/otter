package entity

const (
	// Table role table name
	Table string = "role"

	// PK role table pk column name
	PK string = "code"
	// Code role table code column name
	Code string = "code"
	// Name role table name column name
	Name string = "name"
	// Lv role table lv column name
	Lv string = "lv"
	// SortNo role table sort no column name
	SortNo string = "sort_no"
	// Enable role table enable column name
	Enable string = "enable"
	// CreatedDate role table created date column name
	CreatedDate string = "created_date"
	// UpdatedDate role table updated date column name
	UpdatedDate string = "updated_date"
)

// Role table struct
type Entity struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Lv          int    `json:"lv"`
	SortNo      int    `json:"sortNo"`
	Enable      bool   `json:"enable"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}
