package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

func (rt *_router) sendmessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message_id, err := rt.db.SendMessage(message.Chat_id, message.Owner,
		database.Message{
			message_id: "",
			Chat_id: message.Message_id,
			Status: message.Status,
			Date: message.Date,
			Content: message.Content,
			Forwarded: message.Forwarded,
			Owner: message.Owner,
			Reply: message.Reply,
			Media: message. Media,
		})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	type Resp struct {
		Resp string `json:"chat_id"`
	}
	answ := strconv.FormatInt(message_id, 10)
	_ = json.NewEncoder(w).Encode(Resp{Resp: answ})
}
