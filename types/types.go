package types

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
