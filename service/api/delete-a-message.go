package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
)

func (rt *_router) deleteMessage(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")
	type Owner struct {
		Id string `json:"user_id"`
	}
	var owner Owner
	err := json.NewDecoder(r.Body).Decode(&owner)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	err = rt.db.DeleteMessage(owner.Id, ps.ByName("chat_id"), ps.ByName("message_id"))
	if err != nil {
		// fmt.Println("Error:", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
}
