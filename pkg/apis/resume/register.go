package resume

import (
	"github.com/labstack/echo/v4"
	"github.com/vietanhduong/resume/pkg/utils/env"
	"net/http"
)

type resource struct{}

// Path resume path
// it should be the the same resume output
// path we fetched at start up
var Path string

func init() {
	Path = env.GetEnvAsStringOrFallback("RESUME_OUTPUT_PATH", "resume.yaml")
}

func Register(g *echo.Group) {
	res := resource{}
	g.GET("/", res.home)
}

func (res *resource) home(ctx echo.Context) error {
	svc, err := NewService(Path)
	if err != nil {
		return err
	}
	data := svc.ConvertResumeToResponse()
	return ctx.Render(http.StatusOK, "resume.html", data)
}
