package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)

func (rt *_router) sendmessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message_id,err := rt.db.SendMessage(message.Chat_id, message.Owner, message.Content)
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			//ctx.Logger.WithError(err).Error("session: can't create response json")
			return
		}
	w.WriteHeader(http.StatusOK)
	_ =json.NewEncoder(w).Encode(message_id)
}