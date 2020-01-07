package routes

import (
	"otter/api/codemap"
	"otter/router"
)

// InitCodemapAPI init codemap api
func InitCodemapAPI() {
	groupName := "/codemap"

	// Get
	router.Get(groupName+"/list", codemap.List)

	// Post
	router.Post(groupName, codemap.Update)

	// Put
	router.Put(groupName, codemap.Add)

	// Delete
	router.Delete(groupName, codemap.Delete)
}
