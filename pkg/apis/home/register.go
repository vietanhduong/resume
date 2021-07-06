package home

import (
	"github.com/labstack/echo/v4"
	"net/http"
)

type resource struct{}

func Register(g *echo.Group) {
	res := resource{}
	g.GET("/", res.home)
}

func (res *resource) home(ctx echo.Context) error {
	svc, err := NewService("resume.yaml")
	if err != nil {
		return err
	}
	data := svc.ConvertResumeToResponse()
	return ctx.Render(http.StatusOK, "resume.html", data)
}
