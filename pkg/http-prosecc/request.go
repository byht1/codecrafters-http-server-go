package httpProsecc

import (
	"path"
	"strings"
)

type Request struct {
	Method    string
	Path      string
	Protocol  string
	StaticDir string
	Params    map[string]string
	headers   map[string]string
}

func NewRequest(reqBuffer []byte, staticDir string) Request {
	originalRequest := strings.Split(string(reqBuffer), "\n")
	info := strings.Split(originalRequest[0], " ")

	req := Request{
		Method:    info[0],
		Path:      info[1],
		Protocol:  info[2],
		StaticDir: staticDir,
		Params:    make(map[string]string),
		headers:   make(map[string]string),
	}

	for _, h := range originalRequest[1:] {
		headerLine := strings.Split(h, " ")
		name := strings.ToLower(headerLine[0][:len(headerLine[0])-1])
		req.headers[name] = strings.TrimSpace(strings.Join(headerLine[1:], " "))
	}

	return req
}

func (r *Request) GetHeader(key string) (string, bool) {
	value, isOk := r.headers[strings.ToLower(key)]
	return value, isOk
}

func (r *Request) GetKey() string {
	return path.Join(r.Method, r.Path)
}
