package router

import (
	echo "github.com/labstack/echo"
	middleware "github.com/labstack/echo/middleware"
	// Controller
	controller "template-echo/controller"
)

func CreateApp() *echo.Echo {
	app := echo.New()
	// Set middleware
	app.Use(middleware.Logger())
	app.Use(middleware.Recover())
	// Set health check
	app.GET("/health", controller.HealthCheck)

	return app
}