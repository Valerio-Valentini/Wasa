package api

import (
	"encoding/json"
	"git.sapienzaapps.it/fantasticcoffee/fantastic-coffee-decaffeinated/service/database"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

func (rt *_router) sendmessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	owner := strings.Split(r.Header.Get("Authorization"), " ")[1]

	var message Message
	err := json.NewDecoder(r.Body).Decode(&message)
	if err != nil {
		//fmt.Println("ERRORE IN SEND MESSAGE")
		//fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	message_id, err := rt.db.SendMessage(message.Chat_id, owner,
		database.Message{
			Message_id: "",
			Chat_id:    message.Message_id,
			Status:     0,
			Content:    message.Content,
			Forwarded:  message.Forwarded,
			Owner:      owner,
			Reply:      message.Reply,
		})
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	type Resp struct {
		Resp string `json:"message_id"`
	}
	answ := strconv.FormatInt(message_id, 10)
	_ = json.NewEncoder(w).Encode(Resp{Resp: answ})
}
