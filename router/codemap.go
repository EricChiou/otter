package router

import (
	"otter/api/codemap"

	"github.com/EricChiou/httprouter"
)

func initCodemapAPI() {
	groupName := "/codemap"
	var controller codemap.Controller

	// Get
	// httprouter.Get(groupName+"/list", controller.List)
	get(groupName+"/list", false, nil, controller.List)

	// Post
	httprouter.Post(groupName, controller.Update)

	// Put
	httprouter.Put(groupName, controller.Add)

	// Delete
	httprouter.Delete(groupName, controller.Delete)
}
