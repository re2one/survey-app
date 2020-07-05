package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"survey-app-backend/adapter/controller"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users", func(context echo.Context) error { return c.GetAll(context) })
	e.GET("/users/single", func(context echo.Context) error { return c.Get(context) })
	e.POST("/users", func(context echo.Context) error { return c.Add(context) })
	e.PUT("/users", func(context echo.Context) error { return c.Update(context) })
	e.DELETE("/users", func(context echo.Context) error { return c.Delete(context) })
	return e
}