package httpProsecc

import (
	"path"
)

type Request struct {
	Method   string
	Path     string
	Protocol string
	Params   map[string]string
}

func NewRequest(method, path, protocol string) Request {
	return Request{method, path, protocol, make(map[string]string)}
}

func (r *Request) GetKey() string {
	return path.Join(string(r.Method), r.Path)
}
