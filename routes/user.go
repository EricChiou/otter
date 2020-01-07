package routes

import (
	"otter/api/user"
	"otter/router"
)

// InitUserAPI init user api
func InitUserAPI() {
	groupName := "/user"

	// Get
	router.Get(groupName+"/signIn", user.SignIn)
	router.Get(groupName+"/list", user.List)

	// Post
	router.Post(groupName, user.Update)

	// Put
	router.Put(groupName+"/signUp", user.SignUp)
}
