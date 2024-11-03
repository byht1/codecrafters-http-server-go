package controllers

func NewEndpoint(method, path string, fn RunFunc) *Endpoint {
	return &Endpoint{
		Method: method,
		Path:   path,
		Run:    fn,
	}
}
