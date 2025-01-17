package httpProsecc

import (
	"bytes"
	"compress/gzip"
	"fmt"
	"net"
	"strconv"
	"strings"

	"github.com/codecrafters-io/http-server-starter-go/pkg/compression"
)

const (
	ContentType     = "Content-Type"
	ContentLength   = "Content-Length"
	ContentEncoding = "Content-Encoding"
	AcceptEncoding  = "Accept-Encoding"
)

func gZipData(data []byte) (compressedData []byte, err error) {
	var b bytes.Buffer
	gz := gzip.NewWriter(&b)

	_, err = gz.Write(data)
	if err != nil {
		return
	}

	if err = gz.Flush(); err != nil {
		return
	}

	if err = gz.Close(); err != nil {
		return
	}

	compressedData = b.Bytes()

	return
}

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

	if value, isOk := req.GetHeader(AcceptEncoding); isOk {
		for _, key := range strings.Split(value, ", ") {
			_, ok := compression.Compression[key]
			if ok {
				res.SetHeader(ContentEncoding, key)
				compressedData, err := gZipData(bodyBytes)
				if err == nil {
					bodyBytes = compressedData
				}
			}
		}

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
