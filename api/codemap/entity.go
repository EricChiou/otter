package codemap

// Entity codemap table entity
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

// Col get codemap table column name
func (entity *Entity) Col() Col {
	return Col{
		ID:          "id",
		Type:        "type",
		Code:        "code",
		Name:        "name",
		SortNo:      "sort_no",
		Enable:      "enable",
		CreatedDate: "created_date",
		UpdatedDate: "updated_date",
	}
}

// Col codemap table column name
type Col struct {
	ID          string
	Type        string
	Code        string
	Name        string
	SortNo      string
	Enable      string
	CreatedDate string
	UpdatedDate string
}

// Table codemap table name
func (entity *Entity) Table() string {
	return "codemap"
}

// PK codemap table pk column name
func (entity *Entity) PK() string {
	return "id"
}
