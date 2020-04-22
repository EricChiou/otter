package role

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

// Col get role table column name
func (entity *Entity) Col() Col {
	return Col{
		ID:          "id",
		Code:        "code",
		Name:        "name",
		Lv:          "lv",
		SortNo:      "sort_no",
		Enable:      "enable",
		CreatedDate: "created_date",
		UpdatedDate: "updated_date",
	}
}

// Col role table column name
type Col struct {
	ID          string
	Code        string
	Name        string
	Lv          string
	SortNo      string
	Enable      string
	CreatedDate string
	UpdatedDate string
}

// Table role table name
func (entity *Entity) Table() string {
	return "role"
}

// PK role table pk column name
func (entity *Entity) PK() string {
	return "id"
}
