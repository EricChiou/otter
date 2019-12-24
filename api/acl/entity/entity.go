package entity

// ACL table struct
type ACL struct {
	Col         ACLCol
	Code        string `json:"code"`
	Name        string `json:"name"`
	Type        string `json:"type"`
	Lv          int    `json:"lv"`
	SortNo      int    `json:"sortNo"`
	Enable      bool   `json:"enable"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}

// Table get table name
func (acl *ACL) Table() string {
	return "acl"
}

// ACLCol get acl table column name
type ACLCol struct{}

// PK get table pk column name
func (acl *ACLCol) PK() string {
	return "code"
}

// Code get table Code column name
func (acl *ACLCol) Code() string {
	return "code"
}

// Name get table Name column name
func (acl *ACLCol) Name() string {
	return "name"
}

// Type get table Type column name
func (acl *ACLCol) Type() string {
	return "type"
}

// Lv get table Lv column name
func (acl *ACLCol) Lv() string {
	return "lv"
}

// SortNo get table SortNo column name
func (acl *ACLCol) SortNo() string {
	return "sort_no"
}

// Enable get table Enable column name
func (acl *ACLCol) Enable() string {
	return "enable"
}

// CreatedDate get table CreatedDate column name
func (acl *ACLCol) CreatedDate() string {
	return "created_date"
}

// UpdatedDate get table UpdatedDate column name
func (acl *ACLCol) UpdatedDate() string {
	return "updated_date"
}
