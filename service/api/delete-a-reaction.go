package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext)
{
	w.Header().Set("Content-Type", "application/json")
	var reaction Reaction
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil
	{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := rt.db.SendMessage(reaction.Owner, reaction.Message_id)
	if err != nil
		{
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("session: can't create response json")
			return
		}
	return nil
}