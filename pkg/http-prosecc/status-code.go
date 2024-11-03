package httpProsecc

var messageForStatusCode = make(map[int]string)

func init() {
	messageForStatusCode[200] = "OK"
	messageForStatusCode[201] = "Created"

	messageForStatusCode[400] = "Bad Request"
	messageForStatusCode[401] = "Unauthorized"
	messageForStatusCode[403] = "Forbidden"
	messageForStatusCode[404] = "Not Found"

	messageForStatusCode[500] = "Server Error"
}

func roundDownToNearestHundred(n int) int {
	return (n / 100) * 100
}

func GetMessage(statusCode int) string {
	m, ok := messageForStatusCode[statusCode]
	if !ok {
		return messageForStatusCode[roundDownToNearestHundred(statusCode)]
	}

	return m
}
