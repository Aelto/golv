package golv

import _ "embed"

type View struct {
	Router    Router
	Fragments []Fragment
}

func (view *View) Register() {
	view.Router.Register()

	for _, fragment := range view.Fragments {
		fragment.Register()
	}
}

//go:embed htmx.min.js
var htmx string

func RenderPage() string {
	htmxScriptTag := "<script>" + htmx + "</script>"

	return `
<!doctype html>
<html lang="fr">
  <head>
    <meta charset="utf-8" />
    <meta name="viewport" content="width=device-width" />
    <title>page title</title>
  </head>
  <body>
  	` + htmxScriptTag + `

    %s
  </body>
</html>
	`
}
