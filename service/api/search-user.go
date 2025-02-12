package api

import (
	"encoding/json"
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
)

func (rt *_router) searchUser(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	w.Header().Set("Content-Type", "application/json")

	users, err := rt.db.SearchUser(r.URL.Query().Get("user_id"), strings.Split(r.Header.Get("Authorization"), " ")[1])
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(http.StatusOK)
	_ = json.NewEncoder(w).Encode(users)
}
