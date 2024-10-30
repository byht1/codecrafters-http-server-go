package httpProsecc

import (
	"path"
)

type Method string

type Request struct {
	Method   Method
	Path     string
	Protocol string
}

func NewRequest(method, path, protocol string) Request {
	return Request{Method(method), path, protocol}
}

func (r *Request) GetKey() string {
	return path.Join(string(r.Method), r.Path)
}
