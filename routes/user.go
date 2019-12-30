package routes

import (
	"otter/api/user"
	"otter/router"
)

var groupName string = "/user"

// InitUserAPI init user api
func InitUserAPI() {
	// Get
	router.Get(groupName+"/signIn", user.SignIn)
	router.Get(groupName, user.List)

	// Post
	router.Post(groupName, user.Update)

	// Put
	router.Put(groupName+"/signUp", user.SignUp)
}
