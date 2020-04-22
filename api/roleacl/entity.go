package roleacl

// Entity role_acl table entity
type Entity struct {
	ID          int    `json:"id"`
	RoleCode    string `json:"roleCode"`
	ACLCode     string `json:"aclCode"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}

// Col get role_acl table column name
func (entity *Entity) Col() Col {
	return Col{
		ID:          "id",
		RoleCode:    "role_code",
		ACLCode:     "acl_code",
		CreatedDate: "created_date",
		UpdatedDate: "updated_date",
	}
}

// Col role_acl table column name
type Col struct {
	ID          string
	RoleCode    string
	ACLCode     string
	CreatedDate string
	UpdatedDate string
}

// Table role_acl table name
func (entity *Entity) Table() string {
	return "role_acl"
}

// PK role_acl table pk column name
func (entity *Entity) PK() string {
	return "id"
}
