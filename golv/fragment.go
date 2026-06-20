package golv

type Fragment struct {
	name      string
	endpoints []Endpoint
}

type EndpointBuilder func() []Endpoint

func NewFragment(name string, init EndpointBuilder) Fragment {
	return Fragment{name, init()}
}

func (fragment *Fragment) Register() {
	for _, endpoint := range fragment.endpoints {
		endpoint.RegisterWithPrefix(fragment.name)
	}
}
