package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

func (rt *_router) getMessagesFromChat(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var chat string
	chat = ps.ByName("chat_id")
	messages, err := rt.db.GetMessagesFromChat(chat)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type Response struct {
		Message []database.Message `json:"messages"`
	}
	_ = json.NewEncoder(w).Encode(Response{Message: messages})
}
