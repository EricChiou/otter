package router

import (
	"github.com/valyala/fasthttp"

	cons "otter/constants"
)

// ListenAndServe start http server
func ListenAndServe(port string) error {
	return newFHServer().ListenAndServe(":" + port)
}

// ListenAndServeTLS start https server
func ListenAndServeTLS(port, certPath, keyPath string) error {
	return newFHServer().ListenAndServeTLS(":"+port, certPath, keyPath)
}

func newFHServer() *fasthttp.Server {
	return &fasthttp.Server{
		Name:    cons.ServerName,
		Handler: FasthttpHandler(),
	}
}
