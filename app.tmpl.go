package main

import (
	components "changeme/templates"
	"context"

	"github.com/a-h/templ"
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called at application startup
func (a *App) startup(ctx context.Context) {
	// Perform your setup here
	a.ctx = ctx
}

// domReady is called after front-end resources have been loaded
func (a App) domReady(ctx context.Context) {
	// Add your action here
}

// beforeClose is called when the application is about to quit,
// either by clicking the window close button or calling runtime.Quit.
// Returning true will cause the application to continue, false will continue shutdown as normal.
func (a *App) beforeClose(ctx context.Context) (prevent bool) {
	return false
}

// shutdown is called at application termination
func (a *App) shutdown(ctx context.Context) {
	// Perform your teardown here
}

func NewChiRouter() *chi.Mux {
	r := chi.NewRouter()
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)
	r.Get("/initial", templ.Handler(components.Pages(
		components.IndexForm{Pages: []components.Page{{Label: "Greet Form", Path: "/greet"}},
			Version: components.AppVersion{Version: "0.0.1", UpdateText: "n"}},
	)).ServeHTTP)
	r.Get("/greet", templ.Handler(components.GreetForm("/greet", "#result", "Submit")).ServeHTTP)
	r.Post("/greet", components.Greet)
	return r
}

type Page struct {
	Label string
	Path  string
}
type AppVersion struct {
	Version    string
	UpdateText string
}
type IndexForm struct {
	Pages   []Page
	Version AppVersion
}
