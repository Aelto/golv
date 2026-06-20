package fragments

import (
	"context"
	"fmt"
	"golv/golv"
	"net/http"
)

var FrgCounter = golv.NewFragment("Counter", func() []golv.Endpoint {
	var postIncrease = golv.NewEndpoint("POST /increase", func(w http.ResponseWriter, r *http.Request) {
		type Form struct {
			Number int `form:"Number"`
		}

		form, err := golv.DeserializeForm[Form](r)
		if err != nil {
			fmt.Fprintln(w, "invalid form data")
			return
		}

		CounterRender(form.Number+1).Render(context.Background(), w)
	})

	var postDecrease = golv.NewEndpoint("POST /decrease", func(w http.ResponseWriter, r *http.Request) {
		type Form struct {
			Number int `form:"Number"`
		}

		form, err := golv.DeserializeForm[Form](r)
		if err != nil {
			fmt.Fprintln(w, "invalid form data")
			return
		}

		CounterRender(form.Number-1).Render(context.Background(), w)
	})

	return []golv.Endpoint{postIncrease, postDecrease}
})
