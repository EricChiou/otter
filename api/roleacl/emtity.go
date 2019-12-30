package roleacl

// Table role_acl table name
const Table string = "role_acl"

// Col role_acl table column name
var Col col = col{
	PK:          []string{"role_code", "acl_code"},
	RoleCode:    "role_code",
	ACLCode:     "acl_code",
	CreatedDate: "created_date",
	UpdatedDate: "updated_date",
}

// Entity role_acl table entity
type Entity struct {
	RoleCode    string `json:"roleCode"`
	ACLCode     string `json:"aclCode"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}

type col struct {
	PK          []string
	RoleCode    string
	ACLCode     string
	CreatedDate string
	UpdatedDate string
}
