package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func (rt *_router) forwardMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	type IncomingChat struct {
		Original_chat_id string `json:"chat_id"`
	}
	var original_chat IncomingChat
	err := json.NewDecoder(r.Body).Decode(&original_chat)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	owner := strings.Split(r.Header.Get("Authorization"), " ")[1]
	id, err := rt.db.ForwardMessage(owner,
		original_chat.Original_chat_id,
		ps.ByName("message_id"),
		ps.ByName("chat_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(id)
}
