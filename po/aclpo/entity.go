package aclpo

// table name
const Table string = "acl"

// pk name
const PK string = "acl.id"

// column name
const (
	ID          string = "acl.id"
	Code        string = "acl.code"
	Name        string = "acl.name"
	Type        string = "acl.type"
	Lv          string = "acl.lv"
	SortNo      string = "acl.sort_no"
	Enable      string = "acl.enable"
	CreatedDate string = "acl.created_date"
	UpdatedDate string = "acl.updated_date"
)

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
