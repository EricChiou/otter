package routes

import (
	"otter/api/user"
	"otter/router"
)

// InitUserAPI init user api
func InitUserAPI() {
	groupName := "/user"
	var con user.Controller

	// Get
	router.Get(groupName+"/signIn", con.SignIn)
	router.Get(groupName+"/list", con.List)

	// Post
	router.Post(groupName, con.Update)

	// Put
	router.Put(groupName+"/signUp", con.SignUp)
}
