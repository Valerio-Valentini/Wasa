package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)

func (rt *_router) getAllChats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) ([]Chat, error) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return nil, err
	}

	chats, err := rt.db.GetChats(user.User_id)
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			//ctx.Logger.WithError(err).Error("session: can't create response json")
			return nil, err
		}
	return chats, nil
}