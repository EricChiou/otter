package rolepo

// table name
const Table string = "role"

// pk name
const PK string = "role.id"

// column name
const (
	ID          string = "role.id"
	Code        string = "role.code"
	Name        string = "role.name"
	Lv          string = "role.lv"
	SortNo      string = "role.sort_no"
	Enable      string = "role.enable"
	CreatedDate string = "role.created_date"
	UpdatedDate string = "role.updated_date"
)

// Entity role table entity
type Entity struct {
	ID          int    `json:"id,omitempty"`
	Code        string `json:"code,omitempty"`
	Name        string `json:"name,omitempty"`
	Lv          int    `json:"lv,omitempty"`
	SortNo      int    `json:"sortNo,omitempty"`
	Enable      bool   `json:"enable,omitempty"`
	CreatedDate string `json:"creatDate,omitempty"`
	UpdatedDate string `json:"updateDate,omitempty"`
}
