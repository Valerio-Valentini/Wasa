package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) addAMember(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	user := ps.ByName("member_id")

	exists, err := rt.VerifyUser(user)
	if err != nil || !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	err = rt.db.AddMember(ps.ByName("chat_id"), user)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
