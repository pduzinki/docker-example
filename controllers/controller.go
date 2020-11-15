package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"go-in-docker-example/models"
)

// MessageController ...
type MessageController struct {
	ms models.MessageService
}

// NewMessageController ...
func NewMessageController(ms models.MessageService) *MessageController {
	return &MessageController{
		ms: ms,
	}
}

// Home ...
func (mc *MessageController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello there!")
}

// ReadMessages ...
func (mc *MessageController) ReadMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := mc.ms.GetAll()
	if err != nil {
		// TODO
		return
	}

	fmt.Fprintln(w, messages)
}

// WriteMessage handles POST /message
func (mc *MessageController) WriteMessage(w http.ResponseWriter, r *http.Request) {
	var message models.Message

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		// TODO
		return
	}

	err = json.Unmarshal(body, &message)
	if err != nil {
		// TODO
		return
	}

	err = mc.ms.Create(&message)
	if err != nil {
		// TODO
		return
	}
}
