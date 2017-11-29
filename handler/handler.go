package handler

import (
	"database/sql"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/spf13/viper"
	"github.com/wwsean08/go2/dao"
)

var router = echo.New()
var goDao dao.RedirectorDAO

func init() {
	router.Pre(middleware.RemoveTrailingSlash())
	router.Use(middleware.Recover())
	router.Use(middleware.Gzip())
}

//SetupHandlers ...
func SetupHandlers(conn *sql.DB) *echo.Echo {
	router.Use(middleware.StaticWithConfig(middleware.StaticConfig{
		Root:   viper.GetString("app_base_dir"),
		Browse: false,
	}))

	goDao = dao.NewPostgresDAO(conn)

	//TODO: Setup renderer
	return router
}
