package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) addAMember(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	err := rt.db.AddMember(ps.ByName("chat_id"), ps.ByName("member_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
