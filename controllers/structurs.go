package controllers

import (
	"net"

	httpProsecc "github.com/codecrafters-io/http-server-starter-go/pkg/http-prosecc"
)

type RunFunc func(conn net.Conn, req *httpProsecc.Request, res *httpProsecc.Response)

type ControllerFunc interface {
	Run(conn net.Conn)
}

type Endpoint struct {
	Method string
	Path   string
	Run    RunFunc
}

type Controller struct {
	basePath  string
	endpoints []Endpoint
}
