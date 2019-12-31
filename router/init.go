package router

import (
	"github.com/valyala/fasthttp"
)

// ListenAndServe start http server
func ListenAndServe(port string) error {
	return fasthttp.ListenAndServe(":"+port, FasthttpHandler())
}

// ListenAndServeTLS start https server
func ListenAndServeTLS(port, certPath, keyPath string) error {
	return fasthttp.ListenAndServeTLS(":"+port, certPath, keyPath, FasthttpHandler())
}
