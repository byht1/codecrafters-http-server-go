package controllers

import (
	"net"

	httpProsecc "github.com/codecrafters-io/http-server-starter-go/pkg/http-prosecc"
)

func NewEndpoint(method, path string, fn func(conn net.Conn)) *Endpoint {
	return &Endpoint{
		Method: httpProsecc.Method(method),
		Path:   path,
		Run:    fn,
	}
}
