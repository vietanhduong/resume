package cmd

import (
	"github.com/vietanhduong/resume/pkg/apis/resume"
	"github.com/vietanhduong/resume/pkg/utils/output"
)

// ValidateResume validate input resume path and return exit code
func ValidateResume(path string) int {
	// resume service is contain parse and validate resume path
	_, err := resume.NewService(path)
	if err == nil {
		output.Printf("%s\n", "this file is looking good")
		return 0
	}
	// exit 1 if validate failed
	output.Eprintf("%v\n", err)
	return 1
}
