package controllers

import (
	"net"

	httpProsecc "github.com/codecrafters-io/http-server-starter-go/pkg/http-prosecc"
)

var root = NewEndpoint("GET", "/", func(conn net.Conn, req *httpProsecc.Request, res *httpProsecc.Response) {
	res.StatusCode = 200
})

var echo = NewEndpoint("GET", "/echo/:text", func(conn net.Conn, req *httpProsecc.Request, res *httpProsecc.Response) {
	res.StatusCode = 200
	res.Body = req.Params["text"]
})
