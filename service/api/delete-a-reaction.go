package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) deleteComment(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	// QUI

	owner := Header.Get("Authorization")
	err = rt.db.DeleteReaction(owner, ps.ByName("message_id"), ps.ByName("chat_id"))
	if err != nil {
		// fmt.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
