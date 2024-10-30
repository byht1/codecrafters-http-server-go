package controllers

import (
	"fmt"
	"net"

	"github.com/codecrafters-io/http-server-starter-go/config"
)

var notFound = NewEndpoint(config.NotFound, "", func(conn net.Conn) {
	response := "HTTP/1.1 404 Not Found\r\n\r\n"

	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing response:", err.Error())
		return
	}
})
