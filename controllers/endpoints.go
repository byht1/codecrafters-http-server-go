package controllers

import (
	"fmt"
	"io"
	"net"
	"os"
	"path"
	"strconv"
	"strings"

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

var readFile = NewEndpoint("GET", "/files/:file", func(conn net.Conn, req *httpProsecc.Request, res *httpProsecc.Response) {
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

var createFile = NewEndpoint("POST", "/files/:file", func(conn net.Conn, req *httpProsecc.Request, res *httpProsecc.Response) {
	value, _ := req.GetHeader(httpProsecc.ContentLength)
	size, err := strconv.Atoi(value)
	if err != nil || size <= 0 {
		res.StatusCode = 422
		res.Body = "Invalid data"
		return
	}

	writeFile, err := os.Create(path.Join(req.StaticDir, req.Params["file"]))
	if err != nil {
		fmt.Printf("error when creating an file: %v", err)
		res.StatusCode = 500
		return
	}
	defer writeFile.Close()

	limitedReader := io.LimitReader(strings.NewReader(req.Body), int64(size))
	_, err = io.Copy(writeFile, limitedReader)
	if err != nil {
		fmt.Printf("error when writing data: %v", err)
		res.StatusCode = 500
		return
	}
	res.StatusCode = 201
})
