package golv

import (
	_ "embed"

	"github.com/a-h/templ"
)

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

func RenderPage(child templ.Component) templ.Component {
	htmxScriptTag := "<script>" + htmx + "</script>"

	return renderPage(
		htmxScriptTag,
		child,
	)

}
