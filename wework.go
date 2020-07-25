package acgo

import (
	"bytes"
	"encoding/json"
	"net/http"
)

// SendWeWorkRobotMsg 向企业微信的机器上发送消息
func SendWeWorkRobotMsg(url string, msg string) error {
	type message struct {
		Content string `json:"content"`
	}
	data := struct {
		Msgid   string  `json:"msgid"`
		Msgtype string  `json:"msgtype"`
		Text    message `json:"markdown"`
	}{
		Msgid:   "@all",
		Msgtype: "markdown",
		Text:    message{msg},
	}

	m, err := json.Marshal(data)
	if err != nil {
		return err
	}

	rsp, err := http.Post(url, "application/json", bytes.NewReader(m))
	if err != nil {
		return err
	}

	defer rsp.Body.Close()

	return nil
}
