package resume

import (
	"errors"
	"github.com/labstack/echo/v4"
	"github.com/vietanhduong/resume/pkg/cerrors"
	"github.com/vietanhduong/resume/pkg/github"
	"github.com/vietanhduong/resume/pkg/utils/env"
	"net/http"
)

type resource struct{}

// Path resume path
// it should be the the same resume output
// path we fetched at startup
var Path string

func init() {
	Path = env.GetEnvAsStringOrFallback("RESUME_OUTPUT_PATH", "resume.yaml")
}

func Register(g *echo.Group) {
	res := resource{}
	g.GET("/", res.home)
	g.GET("/fetch", res.fetch)
}

func (res *resource) home(ctx echo.Context) error {
	svc, err := NewService(Path)
	if err != nil {
		return err
	}
	data := svc.ConvertResumeToResponse()
	return ctx.Render(http.StatusOK, "resume.html", data)
}

func (res *resource) fetch(ctx echo.Context) error {
	token := ctx.Request().Header.Get("Authorization")
	if token == "" {
		return cerrors.NewCError(http.StatusUnauthorized, errors.New("unauthorized"))
	}

	// it should be configuration on startup
	user := env.GetEnvAsStringOrFallback("GH_USER", "vietanhduong")
	repo := env.GetEnvAsStringOrFallback("GH_REPO", "resume")
	branch := env.GetEnvAsStringOrFallback("GH_BRANCH", "master")

	gh := github.Github(user, repo, branch)
	// verify permission
	ok, err := gh.CanPush(token)
	if err != nil {
		return err
	}
	if !ok {
		return cerrors.NewCError(http.StatusForbidden, errors.New("permission denied"))
	}
	// fetch resume
	content, err := gh.GetRaw(env.GetEnvAsStringOrFallback("RESUME_PATH", "resume.yaml"))
	if err != nil {
		return err
	}
	// save
	outputPath := env.GetEnvAsStringOrFallback("RESUME_OUTPUT_PATH", "resume.yaml")
	if err := gh.SaveRaw(content, outputPath); err != nil {
		return err
	}
	return ctx.NoContent(http.StatusNoContent)
}
