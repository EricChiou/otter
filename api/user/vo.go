package user

// SignUpReqVo user sign up request data vo
type SignUpReqVo struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
	Name  string `json:"name"`
}

// SignInReqVo user sign in request data vo
type SignInReqVo struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
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

// ListResVo get user list response vo
type ListResVo struct {
	Records []ListDataVo `json:"records"`
	Total   int          `json:"total"`
}

// ListDataVo user list data vo
type ListDataVo struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Identity string `json:"identity"`
	Active   bool   `json:"active"`
}
