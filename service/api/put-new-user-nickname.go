package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)

func (rt *_router) putNewUserNickname(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var user ChangeUserId
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.UpdateUser(user.User_id, user.User_id_2)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//ctx.Logger.WithError(err).Error("session: can't create response json")
		return
	}
}