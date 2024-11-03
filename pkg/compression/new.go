package compression

var Compression map[string]CompressionIF

func init() {
	Compression = make(map[string]CompressionIF)

	Compression["gzip"] = &gzipCompression{}
}
