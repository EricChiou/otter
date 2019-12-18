package routes

import (
	user "otter/api/user/controller"
	"otter/router"
)

var groupName string = "/user"

// InitUserAPI init user api
func InitUserAPI() {
	// Get
	router.Get(groupName+"/signIn", user.SignIn)

	// Post
	router.Post(groupName, user.Update)

	// Put
	router.Put(groupName+"/signUp", user.SignUp)
}
