package user

// Entity user table entity
type Entity struct {
	ID          int    `json:"id,omitempty"`
	Acc         string `json:"acc,omitempty"`
	Pwd         string `json:"pwd,omitempty"`
	Name        string `json:"name,omitempty"`
	RoleCode    string `json:"roleCode,omitempty"`
	Status      string `json:"status,omitempty"`
	CreatedDate string `json:"creatDate,omitempty"`
	UpdatedDate string `json:"updateDate,omitempty"`
}

// Col get entity column name
func (entity *Entity) Col() Col {
	return Col{
		ID:          "id",
		Acc:         "acc",
		Pwd:         "pwd",
		Name:        "name",
		RoleCode:    "role_code",
		Status:      "status",
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
	RoleCode    string
	Status      string
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
