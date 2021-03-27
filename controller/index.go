package controller

import (
	"net/http"
	echo "github.com/labstack/echo"
)

func HealthCheck(ctx echo.Context) error {
	return ctx.String(http.StatusOK, "alive")
}