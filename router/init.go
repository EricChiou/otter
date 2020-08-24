package router

import (
	"otter/config"
	"otter/router/routes"

	"github.com/EricChiou/httprouter"
	"github.com/valyala/fasthttp"
)

// Init init api
func Init() {
	routes.InitUserAPI()
	routes.InitCodemapAPI()
}

// ListenAndServe start http server
func ListenAndServe(port string) error {
	return newFHServer().ListenAndServe(":" + port)
}

// ListenAndServeTLS start https server
func ListenAndServeTLS(port, certPath, keyPath string) error {
	return newFHServer().ListenAndServeTLS(":"+port, certPath, keyPath)
}

// SetHeader add api response header
func SetHeader(key string, value string) {
	httprouter.SetHeader(key, value)
}

func newFHServer() *fasthttp.Server {
	return &fasthttp.Server{
		Name:    config.ServerName,
		Handler: httprouter.FasthttpHandler(),
	}
}
