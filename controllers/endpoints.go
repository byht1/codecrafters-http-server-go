package controllers

import (
	"fmt"
	"net"
)

var root = NewEndpoint("GET", "/", func(conn net.Conn) {
	response := "HTTP/1.1 200 OK\r\n\r\n"

	_, err := conn.Write([]byte(response))
	if err != nil {
		fmt.Println("Error writing response:", err.Error())
		return
	}
})
