package home

import (
	"github.com/labstack/echo/v4"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"net/http"
)

type resource struct{}

func Register(g *echo.Group) {
	res := resource{}
	g.GET("/", res.home)
}

func (res *resource) home(ctx echo.Context) error {
	content, err := ioutil.ReadFile("resume.yaml")
	if err != nil {
		return err
	}

	resume := Resume{}
	err = yaml.Unmarshal(content, &resume)
	if err != nil {
		return err
	}
	err = ValidateResume(resume)
	if err != nil {
		return err
	}
	return ctx.HTML(http.StatusOK, "pass")
}
