package zteMF823Client

import (
	"net/http"
)

func NewClient(host string) *Client {
	return &Client{
		ModemHost:  host,
		httpClient: &http.Client{},
	}
}

type Client struct {
	ModemHost  string
	httpClient *http.Client
}

func (c *Client) SetCustomClient(client *http.Client) {
	c.httpClient = client
}
