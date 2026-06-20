package fragments

import (
	"fmt"
	"golv/golv"
	"net/http"
	"strconv"
)

func FrgCounterRender(value int) string {
	return fmt.Sprintf(`
			<form
				hx-target="this"
				hx-swap="outerHTML"
			>

				<button hx-post="/Counter/decrease">decrease</button>
				<input type="number" name="number" value="%d"></input>
				<button hx-post="/Counter/increase">increase</button>

			</form>
		`, value)
}

var FrgCounter = golv.NewFragment("Counter", func() []golv.Endpoint {
	var postIncrease = golv.NewEndpoint("POST /increase", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintln(w, "invalid form data")
			return
		}

		number, err := strconv.Atoi(r.Form.Get("number"))
		if err != nil {
			fmt.Fprintln(w, "provided value is not a number")
			return
		}

		fmt.Fprint(w, FrgCounterRender(number+1))
	})

	var postDecrease = golv.NewEndpoint("POST /decrease", func(w http.ResponseWriter, r *http.Request) {
		err := r.ParseForm()
		if err != nil {
			fmt.Fprintln(w, "invalid form data")
			return
		}

		number, err := strconv.Atoi(r.Form.Get("number"))
		if err != nil {
			fmt.Fprintln(w, "provided value is not a number")
			return
		}

		fmt.Fprint(w, FrgCounterRender(number-1))
	})

	return []golv.Endpoint{postIncrease, postDecrease}
})
