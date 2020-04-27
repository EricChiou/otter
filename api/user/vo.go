package user

// SignUpReqVo user sign up request data vo
type SignUpReqVo struct {
	Acc  string `json:"acc"`
	Pwd  string `json:"pwd"`
	Name string `json:"name"`
}

// SignInReqVo user sign in request data vo
type SignInReqVo struct {
	Acc string `json:"acc"`
	Pwd string `json:"pwd"`
}

// SignInResVo user sign in response data vo
type SignInResVo struct {
	Token string `json:"token"`
}

// UpdateReqVo update user request data vo
type UpdateReqVo struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
	ID   int    `json:"id"`
}

// ListDataVo user list data vo
type ListDataVo struct {
	ID       int    `json:"id"`
	Acc      string `json:"acc"`
	Name     string `json:"name"`
	RoleCode string `json:"roleCode"`
	Status   string `json:"status"`
}
