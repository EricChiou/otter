package entity

// RoleACL table struct
type RoleACL struct {
	Col         RoleACLCol
	RoleCode    string `json:"roleCode"`
	ACLCode     string `json:"aclCode"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}

// Table get table name
func (roleACL *RoleACL) Table() string {
	return "role_acl"
}

// RoleACLCol get role_acl table column name
type RoleACLCol struct{}

// PK get table pk column name
func (roleacl *RoleACLCol) PK() []string {
	return []string{"role_code", "acl_code"}
}

// RoleCode get table RoleCode column name
func (roleacl *RoleACLCol) RoleCode() string {
	return "role_code"
}

// ACLCode get table ACLCode column name
func (roleacl *RoleACLCol) ACLCode() string {
	return "acl_code"
}

// CreatedDate get table CreatedDate column name
func (roleacl *RoleACLCol) CreatedDate() string {
	return "created_date"
}

// UpdatedDate get table UpdatedDate column name
func (roleacl *RoleACLCol) UpdatedDate() string {
	return "updated_date"
}
