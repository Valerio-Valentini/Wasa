package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) forwardMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	type IncomingChat struct {
		Chat_id string `json:"chats"`
	}
	var original_chat IncomingChat
	err := json.NewDecoder(r.Body).Decode(&original_chat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	owner := r.Header.Get("Authorization")
	id, err := rt.db.ForwardedMessage(owner,
		original_chat.Chat_id,
		ps.ByName("message_id"),
		ps.ByName("chat_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(id)
}
