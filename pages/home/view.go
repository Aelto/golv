package home

import (
	"context"
	"golv/golv"
	"golv/pages/home/fragments"
	"net/http"
)

var ViewHome = golv.View{
	Router: golv.NewRouter(
		[]golv.Endpoint{getIndex},
	),
	Fragments: []golv.Fragment{
		fragments.FrgCounter,
	},
}

var getIndex = golv.NewEndpoint(
	"GET /",
	func(w http.ResponseWriter, r *http.Request) {

		golv.RenderPage(fragments.CounterRender(0)).Render(context.Background(), w)
	},
)
