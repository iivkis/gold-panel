package tgbotcallback

import (
	"fmt"
	"strings"
)

func SplitCallback(data string) (callback string, payload string) {
	s := strings.SplitN(data, "|", 2)

	if len(s) >= 1 {
		callback = s[0]
	}

	if len(s) >= 2 {
		payload = s[1]
	}

	return callback, payload
}

func NewCallbackData(callback string, payload interface{}) string {
	return fmt.Sprintf("%s|%v", callback, payload)
}
