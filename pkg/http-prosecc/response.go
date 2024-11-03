package httpProsecc

type Headers map[string]string

type Response struct {
	headers    Headers
	StatusCode int
	Body       string
}

func NewResponse() Response {
	return Response{
		headers: make(Headers),
	}
}

func (r *Response) GetAllHeaders() Headers {
	return r.headers
}

func (r *Response) GetHeader(name string) (string, bool) {
	value, isOk := r.headers[name]
	return value, isOk
}

func (r *Response) SetHeader(name string, value string) {
	r.headers[name] = value
}
