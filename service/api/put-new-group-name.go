package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strconv"
	"strings"
)

func (rt *_router) putNewGroupName(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	var values NameChat
	err := json.NewDecoder(r.Body).Decode(&values)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	chat_id, err := strconv.Atoi(ps.ByName("chat_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	owner := strings.Split(r.Header.Get("Authorization"), " ")[1]

	err = rt.db.SetGroupName(owner, chat_id, values.Chat_name)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
