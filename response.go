package sfapi

import (
	"net/http"
)

type Response struct {
	Response *http.Response
}

func newResponse(response *http.Response) *Response {
	return &Response{
		Response: response,
	}
}
