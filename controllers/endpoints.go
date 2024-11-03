package controllers

import (
	"fmt"
	"io"
	"net"
	"os"
	"path"

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

var files = NewEndpoint("GET", "/files/:file", func(conn net.Conn, req *httpProsecc.Request, res *httpProsecc.Response) {
	file, err := os.Open(path.Join(req.StaticDir, req.Params["file"]))
	if err != nil {
		fmt.Printf("error opening file: %v\n", err)
		res.StatusCode = 404
		return
	}

	byteData, err := io.ReadAll(file)
	if err != nil {
		fmt.Printf("error reading file: %v\n", err)
		res.StatusCode = 404
		return
	}

	res.StatusCode = 200
	res.File = byteData
})
