package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)	

func (rt *_router) deleteMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message_id,err := rt.db.DeleteMessage(message.Owner, message.Chat_id, message.Message_id)
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			//ctx.Logger.WithError(err).Error("session: can't create response json")
			return
		}
	return 
}