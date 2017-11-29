package handler

import "github.com/labstack/echo"

func init() {
	router.POST("/add", addHandler)
}

func addHandler(context echo.Context) error {

	return nil
}
