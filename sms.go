package zteMF823Client

import (
	"github.com/pkg/errors"
	"net/url"
	"time"
)

type SendSMSRequest struct {
	Text  string
	Phone string
}

type SendSMSResponse struct {
	Result string `json:"result"`
}

func (c *Client) SendSMS(r SendSMSRequest) (*SendSMSResponse, error) {
	respInterface, err := c.Post(&SendSMSResponse{}, url.Values{
		"goformId":    []string{"SEND_SMS"},
		"notCallback": []string{"true"},
		"Number":      []string{r.Phone},
		"sms_time":    []string{time.Now().Format("06;01;02;03;04;05")},
		"MessageBody": []string{c.convertUTF2Unicode(r.Text)},
		"ID":          []string{"-1"},
		"encode_type": []string{"UNICODE"},
	})
	if err != nil {
		return nil, errors.Wrap(err, "send sms request")
	}

	return respInterface.(*SendSMSResponse), nil
}
