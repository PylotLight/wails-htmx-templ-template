package main

import (
	"html/template"
	"net/http"
)

type TextComponent struct {
	Label       string
	Name        string
	Placeholder string
	Classes     string
}

type ButtonComponent struct {
	Label    string
	OnClick  string // Custom onclick event
	HxMethod string // Custom HTMX method (get or post)
	HxURL    string // Custom HTMX URL
	Classes  string
	Target   string //hx-swap target
}

// Define a custom middleware type
type customMiddleware struct {
	Next http.Handler
}

func (c *customMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your custom middleware logic goes here

		if r.URL.Path == "/greet" {
			handleGreetForm(w, r)
			return
		}
		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func handleGreetForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	formData := map[string]interface{}{
		"NameInput": TextComponent{
			Name:        "name",
			Classes:     "",
			Placeholder: "Wails User",
			Label:       "Greet",
		},
		"GreetButton": ButtonComponent{
			Label:    "Greet",
			HxURL:    "/submitgreet",
			HxMethod: "post",
			Target:   "#result",
			// Add other button properties
		},
		// Add other form field data
	}
	tmpl, err := template.New("form").ParseFS(templates, "components/forms.html", "components/inputs.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "greetform", formData)
	if err != nil {
		println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func (a *App) handleGreet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Reswap", "innerHTML")
	w.Write([]byte(a.Greet(r.FormValue("name"))))
}
