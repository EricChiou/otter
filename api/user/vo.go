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
}
