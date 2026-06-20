package fragments

import (
	"fmt"
	"golv/golv"
	"net/http"
)

func FrgCounterRender(value int) string {
	return fmt.Sprintf(`
			<form
				hx-target="this"
				hx-swap="outerHTML"
			>

				<button hx-post="/Counter/decrease">decrease</button>
				<input type="number" name="Number" value="%d"></input>
				<button hx-post="/Counter/increase">increase</button>

			</form>
		`, value)
}

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

		fmt.Fprint(w, FrgCounterRender(form.Number+1))
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

		fmt.Fprint(w, FrgCounterRender(form.Number-1))
	})

	return []golv.Endpoint{postIncrease, postDecrease}
})
