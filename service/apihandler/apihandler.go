package apihandler

import (
	"encoding/json"
	"fmt"
	"otter/api/common"
	"otter/constants/api"

	"github.com/valyala/fasthttp"
)

// ResponseEntity response format
type ResponseEntity struct{}

type apiResponse struct {
	Status api.RespStatus `json:"status"`
	Data   interface{}    `json:"data,omitempty"`
	Trace  interface{}    `json:"trace,omitempty"`
}

var header map[string]string = map[string]string{
	"Content-Type": "application/json",
}

func addHeader(ctx *fasthttp.RequestCtx) {
	for k, v := range header {
		ctx.Response.Header.Add(k, v)
	}
}

// OK api success
func (re *ResponseEntity) OK(ctx *fasthttp.RequestCtx, data interface{}) ResponseEntity {
	addHeader(ctx)

	result := apiResponse{
		Status: api.Success,
		Data:   data,
		Trace:  nil,
	}

	bytes, _ := json.Marshal(result)
	fmt.Fprintf(ctx, string(bytes))
	return *re
}

// Error api error
func (re *ResponseEntity) Error(ctx *fasthttp.RequestCtx, status api.RespStatus, trace interface{}) ResponseEntity {
	addHeader(ctx)

	result := apiResponse{
		Status: status,
		Data:   nil,
		Trace:  trace,
	}

	bytes, _ := json.Marshal(result)
	fmt.Fprintf(ctx, string(bytes))
	return *re
}

// Page api page format
func (re *ResponseEntity) Page(ctx *fasthttp.RequestCtx, list common.PageRespVo, status api.RespStatus, trace interface{}) ResponseEntity {
	addHeader(ctx)

	result := apiResponse{
		Status: status,
		Data:   list,
		Trace:  trace,
	}

	bytes, _ := json.Marshal(result)
	fmt.Fprintf(ctx, string(bytes))
	return *re
}
