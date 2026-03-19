package utils

import (
	"encoding/json"

	"github.com/valyala/fasthttp"
)

func WriteJSONResponse(ctx *fasthttp.RequestCtx, status int, body string) {
	ctx.SetStatusCode(status)
	ctx.SetContentType("application/json")
	ctx.SetBodyString(body)
}

func RespondError(ctx *fasthttp.RequestCtx, status int, message string) {
	WriteJSONResponseObject(ctx, status, map[string]string{"error": message})
}


func WriteJSONResponseObject(ctx *fasthttp.RequestCtx, status int, data any) {
	ctx.SetStatusCode(status)
	ctx.SetContentType("application/json")

	json.NewEncoder(ctx.Response.BodyWriter()).Encode(data)
}
