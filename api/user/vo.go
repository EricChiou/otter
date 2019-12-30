package user

// SignUpReq user sign up request data struct
type SignUpReq struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
	Name  string `json:"name"`
}

// SignInReq user sign in request data struct
type SignInReq struct {
	Email string `json:"email"`
	Pwd   string `json:"pwd"`
}

// SignInRes user sign in response data struct
type SignInRes struct {
	Token string `json:"token"`
}

// UpdateReq update user request data struct
type UpdateReq struct {
	Name string `json:"name"`
	Pwd  string `json:"pwd"`
	ID   int    `json:"id"`
}

// ListRes get user list response
type ListRes struct {
	Records []ListData `json:"records"`
	Total   int        `json:"total"`
}

// ListData user list data struct
type ListData struct {
	ID       int    `json:"id"`
	Email    string `json:"email"`
	Name     string `json:"name"`
	Identity string `json:"identity"`
	Active   bool   `json:"active"`
}
