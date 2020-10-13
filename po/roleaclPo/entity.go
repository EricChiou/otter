package roleaclpo

// table name
const Table string = "role_acl"

// pk name
const PK string = "role_acl.id"

// column name
const (
	ID          string = "role_acl.id"
	RoleCode    string = "role_acl.role_code"
	ACLCode     string = "role_acl.acl_code"
	CreatedDate string = "role_acl.created_date"
	UpdatedDate string = "role_acl.updated_date"
)

// Entity role_acl table entity
type Entity struct {
	ID          int    `json:"id,omitempty"`
	RoleCode    string `json:"roleCode,omitempty"`
	ACLCode     string `json:"aclCode,omitempty"`
	CreatedDate string `json:"creatDate,omitempty"`
	UpdatedDate string `json:"updateDate,omitempty"`
}
