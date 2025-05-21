module changeme

        go 1.18

	require (
	github.com/wailsapp/wails/v2 {{.WailsVersion}}
	github.com/a-h/templ v0.3.865
	github.com/go-chi/chi/v5 v5.2.1
	)

        require (
        github.com/bep/debounce v1.2.1 // indirect
	github.com/go-ole/go-ole v1.3.0 // indirect
	github.com/google/uuid v1.3.1 // indirect
	github.com/jchv/go-winloader v0.0.0-20210711035445-715c2860da7e // indirect
	github.com/labstack/echo/v4 v4.11.2 // indirect
	github.com/labstack/gommon v0.4.0 // indirect
	github.com/leaanthony/go-ansi-parser v1.6.1 // indirect
	github.com/leaanthony/gosod v1.0.3 // indirect
	github.com/leaanthony/slicer v1.6.0 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/pkg/browser v0.0.0-20210911075715-681adbf594b8 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/rivo/uniseg v0.4.4 // indirect
	github.com/samber/lo v1.38.1 // indirect
	github.com/tkrajina/go-reflector v0.5.6 // indirect
	github.com/valyala/bytebufferpool v1.0.0 // indirect
	github.com/valyala/fasttemplate v1.2.2 // indirect
	github.com/wailsapp/go-webview2 v1.0.8 // indirect
	github.com/wailsapp/mimetype v1.4.1 // indirect
	golang.org/x/crypto v0.14.0 // indirect
	golang.org/x/exp v0.0.0-20231006140011-7918f672742d // indirect
	golang.org/x/net v0.17.0 // indirect
	golang.org/x/sys v0.13.0 // indirect
	golang.org/x/text v0.13.0 // indirect
        )

// replace github.com/wailsapp/wails/v2 {{.WailsVersion}} => {{.WailsDirectory}}