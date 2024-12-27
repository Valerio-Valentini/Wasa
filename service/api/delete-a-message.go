package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) deleteMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext)
{
	w.Header().Set("Content-Type", "application/json")
	var message Message
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil
	{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message_id,err := rt.db.SendMessage(message.Owner, message.Chat_id, message.Message_id)
	if err != nil
		{
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("session: can't create response json")
			return
		}
	return nil
}