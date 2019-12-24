package entity

// Role table struct
type Role struct {
	Col         RoleCol
	Code        string `json:"code"`
	Name        string `json:"name"`
	Lv          int    `json:"lv"`
	SortNo      int    `json:"sortNo"`
	Enable      bool   `json:"enable"`
	CreatedDate string `json:"creatDate"`
	UpdatedDate string `json:"updateDate"`
}

// Table get table name
func (role *Role) Table() string {
	return "role"
}

// RoleCol get role table column name
type RoleCol struct{}

// PK get table pk column name
func (acl *RoleCol) PK() string {
	return "code"
}

// Code get table Code column name
func (acl *RoleCol) Code() string {
	return "code"
}

// Name get table Name column name
func (acl *RoleCol) Name() string {
	return "name"
}

// Lv get table Lv column name
func (acl *RoleCol) Lv() string {
	return "lv"
}

// SortNo get table SortNo column name
func (acl *RoleCol) SortNo() string {
	return "sort_no"
}

// Enable get table Enable column name
func (acl *RoleCol) Enable() string {
	return "enable"
}

// CreatedDate get table CreatedDate column name
func (acl *RoleCol) CreatedDate() string {
	return "created_date"
}

// UpdatedDate get table UpdatedDate column name
func (acl *RoleCol) UpdatedDate() string {
	return "updated_date"
}
