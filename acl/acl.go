package acl

import (
	"otter/api/roleacl"
	"otter/db/mysql"
)

// Code acl code type
type Code string

const (
	// AddCodemap acl code
	AddCodemap Code = "addCodemap"
	// UpdateCodemap acl code
	UpdateCodemap Code = "updateCodemap"
	// DeleteCodemap acl code
	DeleteCodemap Code = "deleteCodemap"
	// UpdateUserInfo acl code
	UpdateUser Code = "updateUser"
	// DeleteUser acl code
	DeleteUser Code = "deleteUser"
)

var roleACL map[string][]Code = make(map[string][]Code)

// Load loading permission setting
func Load() error {
	tx, err := mysql.DB.Begin()
	defer tx.Commit()
	if err != nil {
		return err
	}

	// reset roleACL
	roleACL = make(map[string][]Code)

	var entity roleacl.Entity
	column := []string{entity.Col().RoleCode, entity.Col().ACLCode}
	rows, err := mysql.Query(tx, entity.Table(), column, make(map[string]interface{}), "")
	if err != nil {
		return err
	}

	for rows.Next() {
		err = rows.Scan(&entity.RoleCode, &entity.ACLCode)
		if err != nil {
			return err
		}
		if roleACL[entity.RoleCode] == nil {
			roleACL[entity.RoleCode] = []Code{Code(entity.ACLCode)}
		} else {
			roleACL[entity.RoleCode] = append(roleACL[entity.RoleCode], Code(entity.ACLCode))
		}
	}

	return nil
}

// Check check role permission
func Check(aclCode Code, roleCode string) bool {
	if roleACL[roleCode] != nil {
		for _, code := range roleACL[roleCode] {
			if aclCode == code {
				return true
			}
		}
	}
	return false
}
