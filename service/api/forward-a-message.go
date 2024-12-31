package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) sendMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var message ForwardedMessage
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message_id,err := rt.db.ForwardMessage(message.Owner, message.first_chat_id, message.Content, message.second_chat_id)
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("session: can't create response json")
			return
		}
	return message_id, nil
}