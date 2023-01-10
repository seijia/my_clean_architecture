package router

import (
	"api_client/interface/controller"

	"github.com/labstack/echo"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {

	e.Logger.Info("start")
	e.POST("/drm/license", func(context echo.Context) error { return c.User.GetLicense(context) })
	return e
}
