package home

import (
	"github.com/labstack/echo/v4"
	"log"
	"net/http"
)

type resource struct{}

func Register(g *echo.Group) {
	res := resource{}
	g.GET("/", res.home)
}

func (res *resource) home(ctx echo.Context) error {
	service, err := NewService("resume.yaml")
	if err != nil {
		return err
	}

	log.Printf("%+v", service.resume)

	return ctx.HTML(http.StatusOK, "pass")
}
