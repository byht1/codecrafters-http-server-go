package controllers

import (
	"net"

	httpProsecc "github.com/codecrafters-io/http-server-starter-go/pkg/http-prosecc"
)

type ControllerFunc interface {
	Run(conn net.Conn)
}

type Endpoint struct {
	Method httpProsecc.Method
	Path   string
	Run    func(conn net.Conn)
}

type Controller struct {
	basePath  string
	endpoints []Endpoint
}
