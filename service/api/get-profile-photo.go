package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"path/filepath"
)

func (rt *_router) getProfilePhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.ServeFile(w, r, 
	filepath.Join("/tmp/media", ps.ByName("user_id"), "1.jpg"))
}