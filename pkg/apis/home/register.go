package home

import (
	"github.com/labstack/echo/v4"
	"github.com/vietanhduong/resume/pkg/github"
	"net/http"
)

type resource struct{}

func Register(g *echo.Group) {
	res := resource{}
	g.GET("/", res.home)
}

func (res *resource) home(ctx echo.Context) error {
	//service, err := NewService("resume.yaml")
	//if err != nil {
	//	return err
	//}

	githubRaw := github.NewRaw("vietanhduong", "resume", "master")
	resume, err := githubRaw.GetRaw("resume.yaml")
	if err != nil {
		return err
	}
	if err = githubRaw.SaveRaw(resume, "resume.override.yaml"); err != nil {
		return err
	}
	return ctx.HTMLBlob(http.StatusOK, resume)
}
