package health

import (
	"github.com/fasthttp/router"
)

func (h *Handler) RegisterRoutes(r *router.Router) {
	r.GET("/api/health/status", h.CheckHealth)
}