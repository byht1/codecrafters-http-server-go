package controllers

import (
	"net"

	httpProsecc "github.com/codecrafters-io/http-server-starter-go/pkg/http-prosecc"
)

var rootEndpoint = NewEndpoint("GET", "/", func(conn net.Conn, req *httpProsecc.Request, res *httpProsecc.Response) {
	res.StatusCode = 200
})

var echoEndpoint = NewEndpoint("GET", "/echo/:text", func(conn net.Conn, req *httpProsecc.Request, res *httpProsecc.Response) {
	res.StatusCode = 200
	res.Body = req.Params["text"]
})

var userAgentEndpoint = NewEndpoint("GET", "/user-agent", func(conn net.Conn, req *httpProsecc.Request, res *httpProsecc.Response) {
	userAgent, isOk := req.GetHeader("User-Agent")
	if !isOk {
		res.StatusCode = 422
		res.Body = "Invalid user-agent"
	}

	res.StatusCode = 200
	res.Body = userAgent
})
