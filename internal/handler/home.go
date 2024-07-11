package handler

import (
	"github.com/labstack/echo/v4"
	"github.com/waldo2810/collaborator/internal/views"
)

type HomeHandler struct{}

func (h HomeHandler) HandleHome(c echo.Context) error {
    template := views.Home()
    return views.Layout(template, "Collaborator").Render(c.Request().Context(), c.Response())
}
