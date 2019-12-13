package apihandler

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

// Result api result handler
func Result(ctx *fasthttp.RequestCtx, status string, data, trace interface{}) string {
	ctx.Response.Header.Add("Content-Type", "application/json")

	result := apiResult{
		Status: status,
		Data:   data,
		Trace:  trace,
	}

	bytes, _ := json.Marshal(result)
	return string(bytes)
}

type apiResult struct {
	Status string      `json:"status"`
	Data   interface{} `json:"data"`
	Trace  interface{} `json:"trace"`
}
