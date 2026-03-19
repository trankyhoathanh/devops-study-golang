package router

import (
	"devops-study-golang/structure/app"

	"github.com/fasthttp/router"
	"github.com/valyala/fasthttp"
)

func New(a *app.App) fasthttp.RequestHandler {
	r := router.New()

	a.HealthHandler.RegisterRoutes(r)

	r.GET("/health", health)
	
	return SimpleCORS(r.Handler)
}

func health(ctx *fasthttp.RequestCtx) {
	ctx.SetStatusCode(fasthttp.StatusOK)
	ctx.SetContentType("application/json")
	ctx.SetBodyString(`{"status":200,"message":"ok"}`)
}

func SimpleCORS(next fasthttp.RequestHandler) fasthttp.RequestHandler {
	return func(ctx *fasthttp.RequestCtx) {
		ctx.Response.Header.Set("Access-Control-Allow-Origin", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Methods", "*")
		ctx.Response.Header.Set("Access-Control-Allow-Headers", "*")
		
		if string(ctx.Method()) == "OPTIONS" {
			ctx.SetStatusCode(fasthttp.StatusNoContent)
			return
		}
		
		next(ctx)
	}
}