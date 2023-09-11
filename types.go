package main

type TextComponent struct {
	// Label string
	Name        string
	Placeholder string
	Classes     string
	// OnClick  string // Custom onclick event
	// HxMethod string // Custom HTMX method (get or post)
	// HxURL    string // Custom HTMX URL
}

type ButtonComponent struct {
	Label    string
	OnClick  string // Custom onclick event
	HxMethod string // Custom HTMX method (get or post)
	HxURL    string // Custom HTMX URL
	Classes  string
	Target   string //hx-swap target
}
type SelectComponent struct {
	Default  string
	Elements []string
	Classes  string
	// OnClick  string // Custom onclick event
	// HxMethod string // Custom HTMX method (get or post)
	// HxURL    string // Custom HTMX URL
}
type CheckBoxComponent struct {
	Name        string
	Description string
	Checked     string
}
type FormComponent struct {
	Button1 ButtonComponent
	Button2 ButtonComponent
	// Add other form fields as needed
}
