package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func init() {
	router.GET("/", indexHandler)
	router.GET("index", indexHandler)
}

func indexHandler(context echo.Context) error {
	return context.String(http.StatusOK, "Hello World")
}
