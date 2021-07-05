// Package cerrors customize error
package cerrors

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

func HTTPErrorHandler(err error, ctx echo.Context) {
	code := http.StatusInternalServerError
	var msg interface{}
	msg = err.Error()
	if he, ok := err.(*echo.HTTPError); ok {
		code = he.Code
		msg = he.Message
	}

	data := map[string]interface{}{
		"title": msg,
		"code":  code,
	}
	if err := ctx.Render(code, "error.html", data); err != nil {
		ctx.Logger().Error(err)
	}
	ctx.Logger().Error(err)
}
