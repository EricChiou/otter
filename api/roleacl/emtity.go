package roleacl

// PK role_acl table pk column name
var PK []string = []string{"role_code", "acl_code"}

const (
	// Table role_acl table name
	Table string = "role_acl"

	// RoleCode role_acl table role code column name
	RoleCode string = "role_code"
	// ACLCode role_acl table acl code column name
	ACLCode string = "acl_code"
	// CreatedDate role_acl table created date column name
	CreatedDate string = "created_date"
	// UpdatedDate role_acl table updated date column name
	UpdatedDate string = "updated_date"
)

// Entity table struct
type Entity struct {
	RoleCode    string `json:"roleCode"`
	ACLCode     string `json:"aclCode"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}
