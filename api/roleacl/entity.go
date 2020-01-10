package roleacl

// Table role_acl table name
const Table string = "role_acl"

// PK role_acl table pk column name
const PK string = "id"

// Entity role_acl table entity
type Entity struct {
	ID          int    `json:"id"`
	RoleCode    string `json:"roleCode"`
	ACLCode     string `json:"aclCode"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}

// Col role_acl table column name
var Col col = col{
	ID:          "id",
	RoleCode:    "role_code",
	ACLCode:     "acl_code",
	CreatedDate: "created_date",
	UpdatedDate: "updated_date",
}

type col struct {
	ID          string
	RoleCode    string
	ACLCode     string
	CreatedDate string
	UpdatedDate string
}
