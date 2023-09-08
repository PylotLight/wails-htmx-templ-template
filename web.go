package main

import (
	"fmt"
	"net/http"
	"text/template"
)

// Implement the assetserver.Middleware interface
func (m *MyMiddleware) Middleware(next http.Handler) http.Handler {
	fmt.Println("midd")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your custom middleware logic goes here

		if r.URL.Path == "/form1" {
			w.Header().Set("Content-Type", "text/html")
			formData := FormComponent{
				Button1: ButtonComponent{
					Label: "Click Me!",
					HxURL: "/route",
					// Add other button properties
				},
				Button2: ButtonComponent{
					Label: "Form1 Button",
					HxURL: "/form1",
					// Add other button properties
				},
				// Add other form field data
			}
			tmpl, err := template.New("form").ParseFiles("frontend/src/components/form.tmpl")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = tmpl.ExecuteTemplate(w, "form1", formData)
			// err = tmpl.Execute(&buffer, buttonData)
			if err != nil {
				fmt.Println("Error:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		if r.URL.Path == "/form2" {
			w.Header().Set("Content-Type", "text/html")
			formData := FormComponent{
				Button1: ButtonComponent{
					Label: "No Click Me!",
					HxURL: "/form2",
					// Add other button properties
				},
				Button2: ButtonComponent{
					Label: "Form 2 button",
					HxURL: "/form2",
					// Add other button properties
				},
				// Add other form field data
			}
			tmpl, err := template.New("form").ParseFiles("frontend/src/components/form.tmpl")
			if err != nil {
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}

			err = tmpl.ExecuteTemplate(w, "form2", formData)
			// err = tmpl.Execute(&buffer, buttonData)
			if err != nil {
				fmt.Println("Error:", err)
				http.Error(w, err.Error(), http.StatusInternalServerError)
				return
			}
			return
		}
		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
