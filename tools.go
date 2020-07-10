package zteMF823Client

import (
	"fmt"
	"strings"
)

func (c *Client) convertUTF2Unicode(text string) string {
	runes := []rune(text)
	encodeStr := make([]string, len(runes))

	for i, r := range runes {
		encodeStr[i] = fmt.Sprintf("%04X", r)
	}
	return strings.Join(encodeStr, "")
}
