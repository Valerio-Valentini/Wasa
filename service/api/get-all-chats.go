package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
)

func (rt *_router) getAllChats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var user User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chats, err := rt.db.GetChats(user.User_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	type Response struct {
		Chats []database.Chat `json:"chats"`
	}
	if chats == nil {
		chats = []database.Chat{}
	}
	_ = json.NewEncoder(w).Encode(Response{Chats: chats})
}
