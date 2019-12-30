package acl

import (
	"otter/api/roleacl"
	"otter/db/mysql"
)

const (
	// UpdateUserInfo acl code
	UpdateUserInfo string = "updateUserInfo"
	// DeleteUser acl code
	DeleteUser string = "deleteUser"
)

var roleACL map[string][]string = make(map[string][]string)

// Load loading permission setting
func Load() error {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return err
	}

	var entity roleacl.Entity
	column := []string{roleacl.Col.RoleCode, roleacl.Col.ACLCode}
	rows, err := mysql.Query(tx, roleacl.Table, column, make(map[string]interface{}))
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&entity.RoleCode, &entity.ACLCode)
		if err != nil {
			return err
		}
		if roleACL[entity.RoleCode] == nil {
			roleACL[entity.RoleCode] = []string{entity.ACLCode}
		} else {
			roleACL[entity.RoleCode] = append(roleACL[entity.RoleCode], entity.ACLCode)
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
