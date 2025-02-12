package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) session(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// verificare o creare funzione per la verifica dell'esistenza dell'utente
	presence, err := rt.db.VerifyUser(user.User_id)
	if presence {
		// w.WriteHeader(http.StatusOk)
		err = json.NewEncoder(w).Encode(user)
		if err != nil {
			// fmt.Println(err)
			w.WriteHeader(http.StatusInternalServerError)
			// ctx.Logger.WithError(err).Error("session: can't create response json")
			return
		}
		return
	}
	err = rt.db.InsertUser(user.User_id)
	if err != nil {
		// fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		// ctx.Logger.WithError(err).Error("session: can't create response json")
		return
	}
	w.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(w).Encode(user)
	if err != nil {
		// fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
