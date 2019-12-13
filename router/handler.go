package router

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/valyala/fasthttp"
)

func methodHandler(rep http.ResponseWriter, req *http.Request) {
	switch req.Method {
	case http.MethodGet:
		pathHandler(rep, req, trees.Get)
	case http.MethodPost:
		pathHandler(rep, req, trees.Post)
	case http.MethodPut:
		pathHandler(rep, req, trees.Put)
	case http.MethodDelete:
		pathHandler(rep, req, trees.Delete)
	case http.MethodPatch:
		pathHandler(rep, req, trees.Patch)
	case http.MethodHead:
		pathHandler(rep, req, trees.Head)
	case http.MethodOptions:
		pathHandler(rep, req, trees.Options)
	default:
		fmt.Fprintf(rep, "404 page not found")
	}
}

func fasthttpMethodHandler(ctx *fasthttp.RequestCtx) {
	switch string(ctx.Method()) {
	case http.MethodGet:
		fasthttpPathHandler(ctx, trees.Get)
	case http.MethodPost:
		fasthttpPathHandler(ctx, trees.Post)
	case http.MethodPut:
		fasthttpPathHandler(ctx, trees.Put)
	case http.MethodDelete:
		fasthttpPathHandler(ctx, trees.Delete)
	case http.MethodPatch:
		fasthttpPathHandler(ctx, trees.Patch)
	case http.MethodHead:
		fasthttpPathHandler(ctx, trees.Head)
	case http.MethodOptions:
		fasthttpPathHandler(ctx, trees.Options)
	default:
		fmt.Fprintf(ctx, "404 page not found")
	}
}

func pathHandler(rep http.ResponseWriter, req *http.Request, tree *node) {
	params := Params{}
	path := req.RequestURI

	if run := mapping(tree, "", path[1:], &params); run != nil {
		(*run)(&Context{Rep: rep, Req: req, Params: params})
		return
	}

	fmt.Fprintf(rep, "404 page not found")
}

func fasthttpPathHandler(ctx *fasthttp.RequestCtx, tree *node) {
	params := Params{}
	path := strings.SplitN(string(ctx.RequestURI()), "?", 2)[0]

	if run := mapping(tree, "", path[1:], &params); run != nil {
		(*run)(&Context{Ctx: ctx, Params: params})
		return
	}

	fmt.Fprintf(ctx, "404 page not found")
}

func mapping(tree *node, path, pathSeg string, params *Params) *func(*Context) {
	if tree.wildChild {
		*params = append(Params{{Key: tree.path, Value: path}}, *params...)
	}

	if len(pathSeg) == 0 {
		if tree.run != nil {
			return tree.run
		}
		return nil
	}

	path, pathSeg = filterPath(pathSeg)
	for _, child := range tree.children {
		if path == child.path || child.wildChild {
			if run := mapping(child, path, pathSeg, params); run != nil {
				return run
			}
		}
	}

	return nil
}
