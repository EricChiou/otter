package routes

import (
	"otter/api/codemap"

	"github.com/EricChiou/httprouter"
)

// InitCodemapAPI init codemap api
func InitCodemapAPI() {
	groupName := "/codemap"
	var controller codemap.Controller

	// Get
	httprouter.Get(groupName+"/list", controller.List)

	// Post
	httprouter.Post(groupName, controller.Update)

	// Put
	httprouter.Put(groupName, controller.Add)

	// Delete
	httprouter.Delete(groupName, controller.Delete)
}
