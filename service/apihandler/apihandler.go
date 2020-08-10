package apihandler

import (
	"encoding/json"
	"otter/constants/api"

	"github.com/valyala/fasthttp"
)

// Result api result handler
func Result(ctx *fasthttp.RequestCtx, status api.RespStatus, data, trace interface{}) string {
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
	Status api.RespStatus `json:"status"`
	Data   interface{}    `json:"data,omitempty"`
	Trace  interface{}    `json:"trace,omitempty"`
}
