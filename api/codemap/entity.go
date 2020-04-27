package codemap

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
