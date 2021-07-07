package cmd

import (
	"fmt"
	"github.com/vietanhduong/resume/pkg/app"
	"github.com/vietanhduong/resume/pkg/utils/env"
)

func StartResumeServer() {
	application := app.App{}
	// initialize routers before run server
	application.Initialize()
	// run server
	application.Run(fmt.Sprintf(":%s", env.GetEnvAsStringOrFallback("PORT", "8080")))
}
