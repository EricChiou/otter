package router

import (
	"otter/acl"
	"otter/api/user"
)

func initUserAPI() {
	groupName := "/user"
	var controller user.Controller

	// Get
	// sign in
	get(groupName+"/signIn", false, nil, controller.SignIn)

	// user list
	get(groupName+"/list", true, nil, controller.List)

	// Post
	post(groupName, true, nil, controller.Update)
	post(groupName+"/:userID", true, []acl.Code{acl.UpdateUser}, controller.UpdateByUserID)

	// Put
	put(groupName+"/signUp", false, nil, controller.SignUp)

}
