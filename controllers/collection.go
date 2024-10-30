package controllers

import (
	"path"
)

var Controllers = map[string]*Endpoint{}

func init() {
	EndpointBuilder(root)

	EndpointBuilder(notFound)
}

func EndpointBuilder(e *Endpoint) {
	key := path.Join(string(e.Method), e.Path)
	Controllers[key] = e
}
