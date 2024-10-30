package main

import (
	"fmt"
	"net"
	"os"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/config"
	"github.com/codecrafters-io/http-server-starter-go/controllers"
	httpProsecc "github.com/codecrafters-io/http-server-starter-go/pkg/http-prosecc"
)

func main() {
	fmt.Println("Logs from your program will appear here!")

	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		fmt.Println("Failed to bind to port 4221")
		os.Exit(1)
	}

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn)
	}

}

func handleConnection(conn net.Conn) {
	defer conn.Close()

	buffer := make([]byte, 1024)
	_, err := conn.Read(buffer)
	if err != nil {
		fmt.Println("Error reading from connection:", err.Error())
		return
	}

	originalRequest := strings.Split(string(buffer), "\n")
	info := strings.Split(originalRequest[0], " ")

	request := httpProsecc.NewRequest(info[0], info[1], info[2])

	endpoint, isOk := controllers.Controllers[request.GetKey()]
	if !isOk {
		controllers.Controllers[config.NotFound].Run(conn)
		return
	}

	endpoint.Run(conn)

}
