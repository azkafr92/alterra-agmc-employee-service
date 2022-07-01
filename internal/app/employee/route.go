package employee

import (
	"github.com/labstack/echo/v4"
)

func (h *handler) Route(g *echo.Group) {
	g.GET("", h.Get)
	g.GET("/:id", h.GetByID)
}