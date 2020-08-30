package router

import (
	"otter/api/user"

	"github.com/EricChiou/httprouter"
)

func initUserAPI() {
	groupName := "/user"
	var controller user.Controller

	// Get
	httprouter.Get(groupName+"/signIn", controller.SignIn)
	httprouter.Get(groupName+"/list", controller.List)

	// Post
	httprouter.Post(groupName, controller.Update)

	// Put
	httprouter.Put(groupName+"/signUp", controller.SignUp)
}
