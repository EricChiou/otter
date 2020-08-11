package router

import (
	"net/http"

	"github.com/valyala/fasthttp"
)

// Context is use to pass variables between middleware
type Context struct {
	Rep    http.ResponseWriter
	Req    *http.Request
	Ctx    *fasthttp.RequestCtx
	Params Params
}

// GetPathParam get path param
func (context Context) GetPathParam(key string) (string, bool) {
	for _, param := range context.Params {
		if param.Key == key {
			return param.Value, true
		}
	}
	return "", false
}

// Param Params data struct
type Param struct {
	Key, Value string
}

// Params Context Params data struct
type Params []Param

// Trees request node tree
type Trees struct {
	Get, Post, Put, Delete, Patch, Copy, Head, Options *node
}

type header struct {
	key   string
	value string
}

type node struct {
	path      string
	wildChild bool
	run       *func(*Context)
	children  []*node
}
