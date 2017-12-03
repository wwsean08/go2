package handler

import (
	"github.com/labstack/echo"
	"net/http"
)

func init() {
	router.POST("/add", addPostHandler)
	router.GET("/add", addGetHandler)
}

func addGetHandler(context echo.Context) error {

	return nil
}

func addPostHandler(context echo.Context) error {
	kw := context.FormValue("keyword")
	url := context.FormValue("url")
	title := context.FormValue("title")

	kwid, err := goDao.AddKeyword(kw, false)
	if err != nil {
		context.Error(err)
		return err
	}

	urlid, err := goDao.AddURL(url, title)
	if err != nil {
		context.Error(err)
		return err
	}

	err = goDao.AssociateKeywordURL(kwid, urlid)
	if err != nil {
		context.Error(err)
		return err
	}

	return context.String(http.StatusOK, "Insert Successful")
}
