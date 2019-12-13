package routes

import (
	user "otter/api/user/delivery"
	"otter/router"
)

// InitUserAPI init user api
func InitUserAPI() {
	// Get
	router.Get("/user/signIn", user.SignIn)

	// Post
	router.Post("/user", user.Update)

	// Put
	router.Put("/user/signUp", user.SignUp)
}
