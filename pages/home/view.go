package home

import (
	"fmt"
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
		fmt.Fprintf(w, golv.RenderPage(), fragments.FrgCounterRender(0))
	},
)
