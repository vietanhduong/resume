package app

import (
	"context"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/vietanhduong/resume/pkg/apis/resume"
	"github.com/vietanhduong/resume/pkg/cerrors"
	"github.com/vietanhduong/resume/pkg/templates"
	"html/template"
	"log"
	"net/http"
	"os"
	"os/signal"
	"time"
)

type App struct {
	Echo *echo.Echo
}

func (a *App) Initialize() {
	a.Echo = echo.New()
	// configure server
	a.Echo.Use(middleware.Gzip())
	a.Echo.Use(middleware.RemoveTrailingSlash())
	a.Echo.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAcceptEncoding},
		AllowMethods: []string{http.MethodGet, http.MethodPost, http.MethodDelete, http.MethodPut, http.MethodOptions, http.MethodPatch},
	}))
	// static files
	a.Echo.Static("/", "public")
	// register template
	a.Echo.Renderer = &templates.Template{
		Templates: template.Must(template.ParseGlob("public/templates/*.html")),
	}
	// register error handler
	a.Echo.HTTPErrorHandler = cerrors.HTTPErrorHandler
	// customize request log
	a.Echo.Use(middleware.LoggerWithConfig(middleware.LoggerConfig{
		Format: "INFO  | ${time_rfc3339} | ${status} | ${method} ${uri} \n",
		Output: a.Echo.Logger.Output(),
	}))
	// custom logger
	a.Echo.Logger.SetHeader("${level} | ${time_rfc3339} | ${short_file}:${line} | ${message}")

	// register routers
	a.initializeRoutes()
}

func (a *App) Run(addr string) {
	var wait time.Duration
	// make sure you call  `Initialize` before run
	server := &http.Server{
		Addr:         addr,
		WriteTimeout: 15 * time.Second,
		ReadTimeout:  15 * time.Second,
		IdleTimeout:  60 * time.Second,
		Handler:      a.Echo,
	}
	// run the server in a goroutine so that it doesn't block
	go func() {
		log.Printf("server is starting at addr: %s", addr)
		if err := server.ListenAndServe(); err != nil {
			log.Println(err)
		}
	}()

	c := make(chan os.Signal, 1)
	// accept graceful shutdown when quit via SIGINT (ctrl+C)
	// SIGKILL, SIGQUIT or SIGTERM (ctrl+/) will not be caught
	signal.Notify(c, os.Interrupt)
	// block until receive signal
	<-c
	// create a deadline wait for
	ctx, cancel := context.WithTimeout(context.Background(), wait)
	defer cancel()
	// doesn't block if no connections, but will otherwise wait
	// until the timeout deadline
	if err := server.Shutdown(ctx); err != nil {
		log.Printf("error when shutdown server with error: %+v", err)
		return
	}
	log.Println("shutting down")
	os.Exit(0)

}

func (a *App) initializeRoutes() {
	g := a.Echo.Group("")

	// home router
	resume.Register(g)

}
