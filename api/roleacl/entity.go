package roleacl

import (
	"otter/api/common"
)

// Entity role_acl table entity
type Entity struct {
	common.BaseEntity
	ID          int    `json:"id" db:"id"`
	RoleCode    string `json:"roleCode" db:"role_code"`
	ACLCode     string `json:"aclCode" db:"acl_code"`
	CreatedDate string `json:"creatDate" db:"created_date"`
	UpdatedDate string `json:"updateDate" db:"updated_date"`
}

// Table role_acl table name
func (entity *Entity) Table() string {
	return "role_acl"
}

// PK role_acl table pk column name
func (entity *Entity) PK() string {
	return "id"
}
