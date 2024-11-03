package httpProsecc

import (
	"fmt"
	"net"
	"strconv"
	"strings"
)

const (
	ContentType   = "Content-Type"
	ContentLength = "Content-Length"
)

func BuilderResponse(conn net.Conn, req *Request, res *Response) {
	protocol := getProtocol(res.StatusCode)
	bodyBytes := []byte(res.Body)
	var headers []string

	if _, isOk := res.GetHeader(ContentType); !isOk {
		res.SetHeader(ContentType, "text/plain")
	}

	if len(res.File) != 0 {
		res.SetHeader(ContentType, "application/octet-stream")
		bodyBytes = res.File
	}

	res.SetHeader(ContentLength, strconv.Itoa(len(bodyBytes)))
	for key, value := range res.GetAllHeaders() {
		headers = append(headers, fmt.Sprintf("%v: %v", key, value))
	}
	headers = append(headers, "\r\n")

	headersString := strings.Join(headers, "\r\n")

	_, err := conn.Write([]byte(protocol))
	_, err = conn.Write([]byte(headersString))
	_, err = conn.Write(bodyBytes)
	if err != nil {
		fmt.Println("Error writing response:", err.Error())
		return
	}

}

func getProtocol(statusCode int) string {
	return fmt.Sprintf("HTTP/1.1 %v %v\r\n", statusCode, GetMessage(statusCode))
}
