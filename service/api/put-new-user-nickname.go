package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
	"fmt"
) 

func (rt *_router) putNewUserNickname(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	user := strings.Split(r.Header.Get("Authorization"), " ")[1]
	exists, err := rt.VerifyUser(user)
	if err != nil || !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}


	pathUser := ps.ByName("user_id")
	if(user != pathUser) {
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	type newNick struct {
		User_id   string `json:"user_id"`
	}
	var nick newNick
	err = json.NewDecoder(r.Body).Decode(&nick)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}


	err = rt.db.UpdateUser(user, nick.User_id)
	if err != nil {
		fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	/*
		// var user ChangeUserId
	// MODIFICA PRENDERE USER DA HEADER
	// err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}
	// AGGIUNGERE CONTROLLO SE UTENTE ESISTE
	err = rt.db.UpdateUser(user.User_id, user.User_id_2)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// ctx.Logger.WithError(err).Error("session: can't create response json")
		return
	}
		*/
}
