package main

import (
	"fmt"
	"net/http"

	// "text/template"
	"html/template"
)

// Implement the assetserver.Middleware interface
func (m *MyMiddleware) Middleware(next http.Handler) http.Handler {
	fmt.Println("midd")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your custom middleware logic goes here

		if r.URL.Path == "/form1" {
			handleForm1(w, r)
			return
		}
		if r.URL.Path == "/form2" {
			handleForm2(w, r)
			return
		}
		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func handleForm1(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	formData := map[string]interface{}{
		"TitleInput": TextComponent{
			Name:        "Title",
			Classes:     "my-2",
			Placeholder: "Title",
		},
		"ParentInput": TextComponent{
			Name:        "Parent",
			Placeholder: "Parent",
		},
		"Parent2": TextComponent{
			Name:        "Parent2",
			Placeholder: "Parent2",
		},
		"Parent3": TextComponent{
			Name:        "Parent3",
			Placeholder: "Parent3",
		},
		"SelectComponent": SelectComponent{
			Default:  "Select an option",
			Elements: []string{"Item 1", "Item 2"},
			Classes:  "items-center",
		},
		"Submit": ButtonComponent{
			Label:    "Submit",
			HxURL:    "/submitForm1",
			HxMethod: "post",
			Classes:  "my-8",
			// Add other button properties
		},
		// Add other form field data
	}
	tmpl, err := template.New("form").ParseFiles("frontend/src/components/forms.html", "frontend/src/components/inputs.html")
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
}

func handleForm2(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	formData := map[string]interface{}{
		"TextInput": TextComponent{
			Name: "test",
		},
		"Button1": ButtonComponent{
			Label: "No Click Me!",
			HxURL: "/form2",
			// Add other button properties
		},
		// "Button2": ButtonComponent{
		// 	Label: "Form 2 button",
		// 	HxURL: "/form2",
		// 	// Add other button properties
		// },
		// Add other form field data
	}
	tmpl, err := template.New("form").ParseFiles("frontend/src/components/forms.html", "frontend/src/components/inputs.html")
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
}
