package handler

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
)

var router = echo.New()

func init() {
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.Recover())
	router.Use(middleware.Gzip())
}

//SetupHandlers ...
func SetupHandlers() *echo.Echo {
	router.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   viper.GetString("app_base_dir"),
		Browse: false,
	}))

	//TODO: Setup renderer
	return router
}
