package user

// Entity user table entity
type Entity struct {
	ID          int    `json:"id"`
	Acc         string `json:"acc"`
	Pwd         string `json:"pwd"`
	Name        string `json:"name"`
	Role        string `json:"role"`
	Active      bool   `json:"active"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}

// Col get entity column name
func (entity *Entity) Col() Col {
	return Col{
		ID:          "id",
		Acc:         "acc",
		Pwd:         "pwd",
		Name:        "name",
		Role:        "role",
		Active:      "active",
		CreatedDate: "created_date",
		UpdatedDate: "updated_date",
	}
}

// Col user table column name
type Col struct {
	ID          string
	Acc         string
	Pwd         string
	Name        string
	Role        string
	Active      string
	CreatedDate string
	UpdatedDate string
}

// Table user table name
func (entity *Entity) Table() string {
	return "user"
}

// PK user table pk column name
func (entity *Entity) PK() string {
	return "id"
}
