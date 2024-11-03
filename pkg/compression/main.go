package compression

import "fmt"

type CompressionIF interface {
	compressResponse(data []byte)
}

type gzipCompression struct{}

func (c *gzipCompression) compressResponse(data []byte) {
	fmt.Println("Gzip compressing run")
}
