package webclient

import (
	"net/http"
)

type HttpClient interface {
	Request(url string) (*HttpRequest, error)
	Do(req *http.Request) (*http.Response, error)
}
