package cmd

import (
	"fmt"
	"github.com/vietanhduong/resume/pkg/app"
	"github.com/vietanhduong/resume/pkg/github"
	"github.com/vietanhduong/resume/pkg/utils/env"
	"github.com/vietanhduong/resume/pkg/utils/output"
	"os"
)

func fetchResume() {
	user := env.GetEnvAsStringOrFallback("GH_USER", "vietanhduong")
	repo := env.GetEnvAsStringOrFallback("GH_REPO", "resume")
	branch := env.GetEnvAsStringOrFallback("GH_BRANCH", "master")
	gr := github.NewRaw(user, repo, branch)
	// get resume from GitHub
	content, err := gr.GetRaw(env.GetEnvAsStringOrFallback("RESUME_PATH", "resume.yaml"))
	if err != nil {
		output.Eprintf("%v\n", err)
		os.Exit(1)
	}
	output.Printf("Fetch Resume from GitHub => Done\n")
	// save resume
	outputPath := env.GetEnvAsStringOrFallback("RESUME_OUTPUT_PATH", "resume.yaml")
	if err := gr.SaveRaw(content, outputPath); err != nil {
		output.Eprintf("%v\n", err)
		os.Exit(1)
	}
	output.Printf("Write Resume to disk => Done\n")
	output.Printf("Output path: %s\n", outputPath)
}

func StartResumeServer() {
	// handle fetch resume on start up
	// default always fetch
	if v, _ := env.GetEnvAsIntOrFallback("FETCH_RESUME_ON_START_UP", 1); v == 1 {
		// fetch resume from GitHub before starting the server
		fetchResume()
	}

	// ready for start server
	application := app.App{}
	// initialize routers before run server
	application.Initialize()
	// run server
	application.Run(fmt.Sprintf(":%s", env.GetEnvAsStringOrFallback("PORT", "8080")))
}
