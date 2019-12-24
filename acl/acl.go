package acl

import (
	"otter/api/roleacl/entity"
	"otter/db/mysql"
)

const (
	// UpdateUserInfo acl code
	UpdateUserInfo string = "updateUserInfo"
	// DeleteUser acl code
	DeleteUser string = "deleteUser"
)

var roleACL map[string][]string = make(map[string][]string)

// Init inital permission setting
func Init() error {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return err
	}

	var roleacl entity.RoleACL
	column := []string{roleacl.Col.RoleCode(), roleacl.Col.ACLCode()}
	rows, err := mysql.Query(tx, roleacl.Table(), column, make(map[string]interface{}))
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&roleacl.RoleCode, &roleacl.ACLCode)
		if err != nil {
			return err
		}
		if roleACL[roleacl.RoleCode] == nil {
			roleACL[roleacl.RoleCode] = []string{roleacl.ACLCode}
		} else {
			roleACL[roleacl.RoleCode] = append(roleACL[roleacl.RoleCode], roleacl.ACLCode)
		}
	}

	return nil
}

// Check check role permission
func Check(aclCode, roleCode string) bool {
	if roleACL[roleCode] != nil {
		for _, code := range roleACL[roleCode] {
			if aclCode == code {
				return true
			}
		}
	}
	return false
}
