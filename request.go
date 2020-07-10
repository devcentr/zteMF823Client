package zteMF823Client

import (
	"encoding/json"
	"github.com/pkg/errors"
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
)

func (c *Client) post(response interface{}, params url.Values) (interface{}, error) {
	requestUrl := "http://" + c.ModemHost + "/goform/goform_set_cmd_process"

	req, err := http.NewRequest("POST", requestUrl, strings.NewReader(params.Encode()))
	if err != nil {
		return "", errors.Wrap(err, "create request")
	}
	req.Header.Add("Referer", "http://192.168.0.1/index.html")

	resp, err := c.httpClient.Do(req)
	if err != nil {
		return "", errors.Wrap(err, "request do")
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return "", errors.Wrap(err, "read body")
	}

	if err := json.Unmarshal(body, response); err != nil {
		return nil, errors.Wrap(err, "unmarshal response")
	}

	return response, nil
}
