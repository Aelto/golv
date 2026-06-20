package golv

import (
	"log"
	"net/http"
	"strings"
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
