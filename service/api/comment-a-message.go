package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"errors"
)

func (rt *_router) commentMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	// QUI

	owner := r.Header.Get("Authorization")
	err = rt.db.InsertReaction(owner, ps.ByName("reaction_id"), ps.ByName("message_id", ps.ByName("chat_id")))
	if err != nil {
		// fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}