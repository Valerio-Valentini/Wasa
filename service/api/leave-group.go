package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) leaveChat(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	owner := r.Header.Get("Authorization")
	member_id := ps.ByName("member_id") 		// QUI
	if owner != member_id {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	err := rt.db.LeaveChat(ps.ByName("chat_id"), member_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// ctx.Logger.WithError(err).Error("session: can't create response json")
		return
	}
}
