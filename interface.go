package zteMF823Client

import "net/http"

type HttpClientInterface interface {
	Do(req *http.Request) (*http.Response, error)
}
