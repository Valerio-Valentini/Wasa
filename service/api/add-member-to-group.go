package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"encoding/json"
)

func (rt *_router) addAMember(w http.ResponseWriter, r *http.Request, ps httprouter.Params) error{
	w.Header().Set("Content-Type", "application/json")
	var combo UserChatCombo
	err := json.NewDecoder(r.Body).Decode(&combo)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		return err
	}

	err = rt.db.AddMember(combo.Chat_id, combo.User_id)
	if err != nil{
			w.WriteHeader(http.StatusInternalServerError)
			//ctx.Logger.WithError(err).Error("session: can't create response json")
			return err
		}
	return nil
}