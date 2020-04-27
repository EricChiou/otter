package apihandler

import (
	"encoding/json"
	cons "otter/constants"

	"github.com/valyala/fasthttp"
)

// Result api result handler
func Result(ctx *fasthttp.RequestCtx, status cons.ApiResult, data, trace interface{}) string {
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
	Status cons.ApiResult `json:"status"`
	Data   interface{}    `json:"data,omitempty"`
	Trace  interface{}    `json:"trace,omitempty"`
}
