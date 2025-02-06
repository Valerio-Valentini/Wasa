package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) createNewChat(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	err := rt.db.NewChat(ps.ByName("chat_group"), ps.ByName("photo_id"), ps.ByName("chat_name"))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}