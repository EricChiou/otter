package constants

// UserStatus user status
var UserStstus userStatus = userStatus{
	Active:   "active",
	Inactive: "inactive",
	Blocked:  "blocked",
}

// UserStatus user status type
type UserStatus string
type userStatus struct {
	Active   UserStatus
	Inactive UserStatus
	Blocked  UserStatus
}
