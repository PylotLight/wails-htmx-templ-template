package main

import (
	"html/template"
	"net/http"
)

// Implement the assetserver.Middleware interface
func (m *MyMiddleware) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Your custom middleware logic goes here

		// if r.URL.Path == "/" {
		// 	tmpl, err := template.New("sidebar").ParseFS(templates, "components/main.html", "components/forms.html", "components/inputs.html")
		// 	// fmt.Print(tmpl, err.Error())
		// 	if err != nil {
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}

		// 	err = tmpl.ExecuteTemplate(w, "sidebar", nil)
		// 	if err != nil {
		// 		println("Error:", err)
		// 		http.Error(w, err.Error(), http.StatusInternalServerError)
		// 		return
		// 	}
		// }

		if r.URL.Path == "/normal" {
			handleNormal(w, r)
			return
		}
		if r.URL.Path == "/release" {
			handleReleaseForm(w, r)
			return
		}
		if r.URL.Path == "/releasesubmit" {
			handleRelease(w, r)
			return
		}
		if r.URL.Path == "/greet" {
			handleGreetForm(w, r)
			return
		}
		if r.URL.Path == "/greetsubmit" {
			NewApp().handleGreet(w, r)
			return
		}
		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}

func YourHandler(w http.ResponseWriter, r *http.Request) {
	// Handle your request and render your template
	tmpl, _ := template.New("form").ParseFS(templates, "components/forms.html", "components/inputs.html")
	// Determine which menu item was clicked and pass it to the template
	menuItem := r.FormValue("menu-item")

	// Render your template with the selected menu item
	tmpl.ExecuteTemplate(w, "menu", struct {
		SelectedMenuItem string
	}{SelectedMenuItem: menuItem})
}

func handleGreetForm(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	formData := map[string]interface{}{
		"NameInput": TextComponent{
			Name:        "name",
			Classes:     "",
			Placeholder: "Greet",
		},
		"GreetButton": ButtonComponent{
			Label:    "Greet",
			HxURL:    "/greetsubmit",
			HxMethod: "post",
			Classes:  "join-item",
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

func handleNormal(w http.ResponseWriter, r *http.Request) {
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
		"RegionAU": CheckBoxComponent{
			Name:        "RegionAU",
			Description: "RegionAU",
			Checked:     "checked",
		},
		"RegionEU": CheckBoxComponent{
			Name:        "RegionEU",
			Description: "RegionEU",
			Checked:     "",
		},
		"RegionUS": CheckBoxComponent{
			Name:        "RegionUS",
			Description: "RegionUS",
			Checked:     "",
		},
		"PDFUpload": CheckBoxComponent{
			Name:        "PDFUpload",
			Description: "PDFUpload",
			Checked:     "checked",
		},
		"Attachments": CheckBoxComponent{
			Name:        "Attachments",
			Description: "Upload Attachments",
			Checked:     "",
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
	tmpl, err := template.New("form").ParseFS(templates, "components/forms.html", "components/inputs.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "normal", formData)
	if err != nil {
		println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleReleaseForm(w http.ResponseWriter, r *http.Request) {
	formData := map[string]interface{}{
		"ParentInput": TextComponent{
			Name:        "Parent",
			Placeholder: "RLSE12356",
		},
		"Version": TextComponent{
			Name:        "Version",
			Classes:     "my-2",
			Placeholder: "4.1.0",
		},
		"RegionAU": CheckBoxComponent{
			Name:        "RegionAU",
			Description: "RegionAU",
			Checked:     "checked",
		},
		"RegionEU": CheckBoxComponent{
			Name:        "RegionEU",
			Description: "RegionEU",
			Checked:     "checked",
		},
		"RegionUS": CheckBoxComponent{
			Name:        "RegionUS",
			Description: "RegionUS",
			Checked:     "checked",
		},
		"PDFUpload": CheckBoxComponent{
			Name:        "PDFUpload",
			Description: "PDFUpload",
			Checked:     "checked",
		},
		"Attachments": CheckBoxComponent{
			Name:        "Attachments",
			Description: "Upload Attachments",
			Checked:     "checked",
		},
		"Submit": ButtonComponent{
			Label:    "Submit",
			HxURL:    "/releasesubmit",
			HxMethod: "post",
			Target:   "#toast",
			// Add other button properties
		},
		"Form": r.URL.EscapedPath(),
		// Add other form field data
	}
	println(r.URL.Path)
	tmpl, err := template.New("form").ParseFS(templates, "components/forms.html", "components/inputs.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "release", formData)
	if err != nil {
		println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
}

func handleRelease(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("HX-Reswap", "innerHTML")
	formData := map[string]interface{}{
		"Message": struct {
			Message string
		}{"Completed Successfully"},
		"Form": r.URL.EscapedPath(),
	}

	println(r.URL.Path)
	tmpl, err := template.New("toast").ParseFS(templates, "components/forms.html", "components/inputs.html")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	err = tmpl.ExecuteTemplate(w, "toast-success", formData)
	if err != nil {
		println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	formData = map[string]interface{}{
		"Message": struct {
			Message string
		}{"Completed"},
		"Form": r.URL.EscapedPath(),
	}

	err = tmpl.ExecuteTemplate(w, "toast-info", formData)
	if err != nil {
		println("Error:", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}
