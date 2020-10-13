package codemappo

// table name
const Table string = "codemap"

// pk name
const PK string = "codemap.id"

// column name
const (
	ID          string = "codemap.id"
	Type        string = "codemap.type"
	Code        string = "codemap.code"
	Name        string = "codemap.name"
	SortNo      string = "codemap.sort_no"
	Enable      string = "codemap.enable"
	CreatedDate string = "codemap.created_date"
	UpdatedDate string = "codemap.updated_date"
)

// Entity codemap table entity
type Entity struct {
	ID          int    `json:"id,omitempty"`
	Type        string `json:"type,omitempty"`
	Code        string `json:"code,omitempty"`
	Name        string `json:"name,omitempty"`
	SortNo      int    `json:"sortNo,omitempty"`
	Enable      bool   `json:"enable,omitempty"`
	CreatedDate string `json:"creatDate,omitempty"`
	UpdatedDate string `json:"updateDate,omitempty"`
}
