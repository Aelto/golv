package golv

type Router struct {
	Endpoints []Endpoint
}

func NewRouter(endpoints []Endpoint) Router {
	return Router{
		Endpoints: endpoints,
	}
}

func (router *Router) Register() {
	for _, endpoint := range router.Endpoints {
		endpoint.Register()
	}
}
