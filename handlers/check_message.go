package handlers

import "strings"

type message struct {
	MessageText string `json:"message"`
}

type response struct {
	Message string `json"message"`
	Banned  bool   `json:"banned"`
}

func (m *message) checkMessage() bool {
	m.MessageText = strings.ToLower(m.MessageText)
	if strings.Contains(m.MessageText, "fuck") {
		return true
	}
	return false

}

func CheckMessage(msg string) response {
	return response{
		Message: msg,
		Banned:  false,
	}
}
