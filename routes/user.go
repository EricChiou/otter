package routes

import (
	"otter/api/user"
	"otter/pkg/router"
)

// InitUserAPI init user api
func InitUserAPI() {
	groupName := "/user"
	var controller user.Controller

	// Get
	router.Get(groupName+"/signIn", controller.SignIn)
	router.Get(groupName+"/list", controller.List)

	// Post
	router.Post(groupName, controller.Update)

	// Put
	router.Put(groupName+"/signUp", controller.SignUp)
}
