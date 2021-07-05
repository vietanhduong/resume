package main

import "github.com/vietanhduong/resume/pkg/app"

func main() {
	application := app.App{}
	// initialize routers before run server
	application.Initialize()
	// run server
	application.Run(":8080")
}
