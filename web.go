package main

import (
	"fmt"
	"html/template"
	"net/http"
	"os"
)

// Implement the assetserver.Middleware interface
func (m *MyMiddleware) Middleware(next http.Handler) http.Handler {
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
		if r.URL.Path == "/greet" {
			handleGreet(w, r)
			return
		}
		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func handleGreet(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	// files, _ := os.ReadDir("/")
	files, _ := os.ReadDir(".")
	loc, _ := os.Getwd()
	fmt.Println(files, loc)
	formData := map[string]interface{}{
		"NameInput": TextComponent{
			Name:        "Greet",
			Classes:     "",
			Placeholder: "Greet",
		},
		"GreetButton": ButtonComponent{
			Label:    "Greet",
			HxURL:    "/greetsubmit",
			HxMethod: "get",
			Classes:  "",
			// Add other button properties
		},
		// Add other form field data
	}
	tmpl, err := template.New("form").ParseFiles("components/forms.html", "components/inputs.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "greetform", formData)
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
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
	tmpl, err := template.New("form").ParseFiles("./components/forms.html", "./components/inputs.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "form1", formData)
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
	}
	tmpl, err := template.New("form").ParseFiles("./components/forms.html", "./components/inputs.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "form2", formData)
	if err != nil {
		fmt.Println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}
