package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"docker-example/models"
)

// MessageController is responsible for handling message resources
type MessageController struct {
	ms models.MessageService
}

// NewMessageController creates message controller instance
func NewMessageController(ms models.MessageService) *MessageController {
	return &MessageController{
		ms: ms,
	}
}

// Home handles GET /
func (mc *MessageController) Home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "hello there!")
}

// ReadMessages handles GET /messages
func (mc *MessageController) ReadMessages(w http.ResponseWriter, r *http.Request) {
	messages, err := mc.ms.GetAll()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	fmt.Fprintln(w, messages)
}

// WriteMessage handles POST /message
func (mc *MessageController) WriteMessage(w http.ResponseWriter, r *http.Request) {
	var message models.Message

	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = json.Unmarshal(body, &message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = mc.ms.Create(&message)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
