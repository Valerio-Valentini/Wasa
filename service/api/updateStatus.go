package api

import (
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func (rt *_router) updateStatus(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	user := strings.Split(r.Header.Get("Authorization"), " ")[1]
	exists, err := rt.VerifyUser(user)
	if err != nil || !exists {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	type data struct {
		User_id    string `json:"user_id"`
		Chat_id    string `json:"chat_id"`
		Message_id string `json:"message_id"`
	}
	var Data data
	err = json.NewDecoder(r.Body).Decode(&Data)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		// fmt.Println(err)
		return
	}
	err = rt.db.UpdateMessageStatus(Data.User_id, Data.Chat_id, Data.Message_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}
