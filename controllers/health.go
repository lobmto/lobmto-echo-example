package controllers

import (
	"lobmto-echo-example/responses"
	"net/http"

	"github.com/labstack/echo/v4"
)

type HealthController struct{}

func NewHealthController() *HealthController {
	return &HealthController{}
}

func (h *HealthController) Check(c echo.Context) error {
	return c.JSON(http.StatusOK, responses.NewHealthResponse())
}
