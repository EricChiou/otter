package entity

const (
	// Table acl table name
	Table string = "acl"

	// PK acl table pk column name
	PK string = "code"
	// Code acl table code column name
	Code string = "code"
	// Name acl table name column name
	Name string = "name"
	// Type acl table type column name
	Type string = "type"
	// Lv acl table lv column name
	Lv string = "lv"
	// SortNo acl table sort no column name
	SortNo string = "sort_no"
	// Enable acl table enable column name
	Enable string = "enable"
	// CreatedDate acl table created date column name
	CreatedDate string = "created_date"
	// UpdatedDate acl table updated date column name
	UpdatedDate string = "updated_date"
)

// ACL table struct
type ACL struct {
	Code        string `json:"code"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Lv          int    `json:"lv"`
	SortNo      int    `json:"sortNo"`
	Enable      bool   `json:"enable"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}
