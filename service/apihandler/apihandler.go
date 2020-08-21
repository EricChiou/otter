package apihandler

import (
	"encoding/json"
	"fmt"
	"otter/api/common"
	"otter/constants/api"

	"github.com/valyala/fasthttp"
)

// Response api response handler
func Response(ctx *fasthttp.RequestCtx, status api.RespStatus, data, trace interface{}) {
	ctx.Response.Header.Add("Content-Type", "application/json")

	result := apiResponse{
		Status: status,
		Data:   data,
		Trace:  trace,
	}

	bytes, _ := json.Marshal(result)
	fmt.Fprintf(ctx, string(bytes))
}

// ResponsePage api response page handler
func ResponsePage(ctx *fasthttp.RequestCtx, status api.RespStatus, list common.PageRespVo, trace interface{}) {
	ctx.Response.Header.Add("Content-Type", "application/json")

	result := apiResponse{
		Status: status,
		Data:   list,
		Trace:  trace,
	}

	bytes, _ := json.Marshal(result)
	fmt.Fprintf(ctx, string(bytes))
}

type apiResponse struct {
	Status api.RespStatus `json:"status"`
	Data   interface{}    `json:"data,omitempty"`
	Trace  interface{}    `json:"trace,omitempty"`
}
