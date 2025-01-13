package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var reaction Reaction
	err := json.NewDecoder(r.Body).Decode(&reaction)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.DeleteReaction(reaction.Owner, reaction.Message_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// ctx.Logger.WithError(err).Error("session: can't create response json")
		return
	}
}
