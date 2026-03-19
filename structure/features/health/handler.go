package health

import (
	"devops-study-golang/pkg"
	"devops-study-golang/structure/utils"
	"time"

	"github.com/valyala/fasthttp"
)

type Handler struct {
	cfg *pkg.Config
}

func NewHandler(cfg *pkg.Config) *Handler {
	return &Handler{
		cfg: cfg,
	}
}

// Health check handler
func (h *Handler) CheckHealth(ctx *fasthttp.RequestCtx) {
	utils.WriteJSONResponse(
		ctx,
		fasthttp.StatusOK,
		`{"status":"Check Health Function ` + h.cfg.ENV + `","timestamp":"` + time.Now().Format(time.RFC3339) + `"}`,
	)
}