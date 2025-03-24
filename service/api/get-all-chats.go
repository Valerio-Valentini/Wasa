package api

import (
	"encoding/json"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getAllChats(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	chats, err := rt.db.GetChats(ps.ByName("user_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	type Response struct {
		Chats []database.ChatPreview `json:"chats"`
	}
	if chats == nil {
		chats = []database.ChatPreview{}
	}
	_ = json.NewEncoder(w).Encode(Response{Chats: chats})
}
