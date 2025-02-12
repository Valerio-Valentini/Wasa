package api

import (
	"encoding/json"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) getMessagesFromChat(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	messages, err := rt.db.GetMessagesFromChat(ps.ByName("chat_id"))
	if err != nil {
		// fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type Response struct {
		Messages []database.Message `json:"messages"`
	}
	_ = json.NewEncoder(w).Encode(Response{Messages: messages})
}
