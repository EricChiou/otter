package acl

import (
	"otter/api/roleacl"
	"otter/db/mysql"
)

const (
	// AddCodemap acl code
	AddCodemap string = "addCodemap"
	// UpdateCodemap acl code
	UpdateCodemap string = "updateCodemap"
	// DeleteCodemap acl code
	DeleteCodemap string = "deleteCodemap"
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

	// reset roleACL
	roleACL = make(map[string][]string)

	var entity roleacl.Entity
	column := []string{entity.Col().RoleCode, entity.Col().ACLCode}
	rows, err := mysql.Query(tx, entity.Table(), column, make(map[string]interface{}))
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
