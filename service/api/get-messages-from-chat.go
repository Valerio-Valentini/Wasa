package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getMessagesFromChat(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext)
{
	w.Header().Set("Content-Type", "application/json")
	var chat int
	err := json.NewDecoder(r.Body).Decode(&chat)
	if err != nil
	{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	messages,err := rt.db.GetChats(chat)
	if err != nil
		{
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("session: can't create response json")
			return
		}
	return messages, nil
}