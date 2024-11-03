package main

import (
	"fmt"
	"net"
	"os"
	"strings"

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

	trie := controllers.NewTrie()

	for {
		conn, err := l.Accept()
		if err != nil {
			fmt.Println("Error accepting connection: ", err.Error())
			os.Exit(1)
		}

		go handleConnection(conn, *trie)
	}

}

func handleConnection(conn net.Conn, tree controllers.Trie) {
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
	response := httpProsecc.NewResponse()

	endpoint, params := tree.FindRoute(request.Method, request.Path)
	if endpoint != nil {
		request.Params = params
		endpoint.Run(conn, &request, &response)
	} else {
		response.StatusCode = 404
	}

	httpProsecc.BuilderResponse(conn, &request, &response)
}
