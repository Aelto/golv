package golv

import (
	"log"
	"net/http"
	"strings"

	"github.com/go-playground/form"
)

type Endpoint struct {
	route   string
	handler http.HandlerFunc
}

func NewEndpoint(route string, handler http.HandlerFunc) Endpoint {
	return Endpoint{route, handler}
}

func (endpoint *Endpoint) Register() {
	http.HandleFunc(endpoint.route, endpoint.handler)
}

func (endpoint *Endpoint) RegisterWithPrefix(prefix_without_slash string) {
	var slices = strings.Split(endpoint.route, " ")
	var route = strings.Join(slices, " /"+prefix_without_slash)

	log.Printf("registering route: %s", route)

	http.HandleFunc(route, endpoint.handler)
}

// use a single instance of Decoder, it caches struct info
var decoder *form.Decoder

func DeserializeForm[DES any](r *http.Request) (DES, error) {
	var output DES
	decoder = form.NewDecoder()

	// this simulates the results of http.Request's ParseForm() function
	err := r.ParseForm()
	if err != nil {
		return output, err
	}

	// must pass a pointer
	err = decoder.Decode(&output, r.Form)
	if err != nil {
		return output, err
	}

	return output, nil
}
