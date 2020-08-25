package acl

import (
	"database/sql"

	"otter/api/roleacl"
	"otter/db/mysql"
)

var DB *sql.DB

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

func test(rows *sql.Rows) {

}

// Load loading permission setting
func Load() error {
	// reset roleACL
	roleACL = make(map[string][]Code)

	var entity roleacl.Entity
	sql := "SELECT #roleCode, #aclCode FROM #roleAclT"
	param := mysql.SQLParamsInstance()
	param.Add("roleAclT", entity.Table())
	param.Add("roleCode", entity.Col().RoleCode)
	param.Add("aclCode", entity.Col().ACLCode)

	return mysql.Query(sql, param, func(result mysql.Rows) error {
		rows := result.Rows

		for rows.Next() {
			err := rows.Scan(&entity.RoleCode, &entity.ACLCode)
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
	})
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
