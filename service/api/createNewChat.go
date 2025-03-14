package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func (rt *_router) createNewChat(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

	type sender struct {
		User_ids   []string `json:"user_ids"`
		Chat_group bool     `json:"chat_group"`
	}

	var data sender

	err := json.NewDecoder(r.Body).Decode(&data)
	if err != nil {
		// fmt.Println(err)
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chat_id, err := rt.db.StartChat(data.Chat_group, data.User_ids)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	id_string := strconv.FormatInt(chat_id, 10)
	err = os.MkdirAll(filepath.Join("/tmp/media", id_string, "photo"), os.ModePerm)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	type response struct {
		Chat_id int64 `json:"chat_id"`
	}
	_ = json.NewEncoder(w).Encode(response{Chat_id: chat_id})
}
