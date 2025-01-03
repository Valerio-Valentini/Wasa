package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)

func (rt *_router) forwardMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var message ForwardedMessage
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	id, err := rt.db.ForwardMessage(message.Owner, message.Chat_id, message.Content, message.Chat_id_2)
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			//ctx.Logger.WithError(err).Error("session: can't create response json")
			return
		}
	w.WriteHeader(http.StatusOK)
	_ =json.NewEncoder(w).Encode(id)
}