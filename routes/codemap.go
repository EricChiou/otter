package routes

import (
	"otter/api/codemap"
	"otter/router"
)

// InitCodemapAPI init codemap api
func InitCodemapAPI() {
	groupName := "/codemap"
	var controller codemap.Controller

	// Get
	router.Get(groupName+"/list", controller.List)

	// Post
	router.Post(groupName, controller.Update)

	// Put
	router.Put(groupName, controller.Add)

	// Delete
	router.Delete(groupName, controller.Delete)
}
