package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)

func (rt *_router) addMember(w http.ResponseWriter, r *http.Request, ps httprouter.Params){
	w.Header().Set("Content-Type", "application/json")
	var combo UserChatCombo
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err := rt.db.AddMember(combo.Chat_id, combo.User_id)
	if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("session: can't create response json")
			return
		}
	return nil
}