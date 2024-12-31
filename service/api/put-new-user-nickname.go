package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) putNewUserNickname(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var user ChangeUserId
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := rt.db.UpdateUser(user.old_user_id, user.new_user_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("session: can't create response json")
		return
	}
	return
}