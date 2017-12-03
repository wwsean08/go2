// Copyright © 2017 Sean Smith <sean@wwsean08.com>
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

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
