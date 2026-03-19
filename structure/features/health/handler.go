package health

import (
	"devops-study-golang/structure/utils"
	"time"

	"github.com/valyala/fasthttp"
)

type Handler struct {}

func NewHandler() *Handler {
	return &Handler{}
}

// Health check handler
func (h *Handler) CheckHealth(ctx *fasthttp.RequestCtx) {
	utils.WriteJSONResponse(
		ctx,
		fasthttp.StatusOK,
		`{"status":"Check Health Function","timestamp":"` + time.Now().Format(time.RFC3339) + `"}`,
	)
}