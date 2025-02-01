package api

import (
	"encoding/json"
	"net/http"
	"github.com/julienschmidt/httprouter"
)

func (rt *_router) putNewGroupName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var values NameChat
	err := json.NewDecoder(r.Body).Decode(&values)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.SetGroupName(values.User_id, values.Chat_id, values.Chat_name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	} 
}
