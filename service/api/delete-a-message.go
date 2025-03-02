package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func (rt *_router) deleteMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	owner := strings.Split(r.Header.Get("Authorization"), " ")[1]

	err := rt.db.DeleteMessage(owner, ps.ByName("chat_id"), ps.ByName("message_id"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
