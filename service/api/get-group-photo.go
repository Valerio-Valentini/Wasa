package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"path/filepath"
)

func (rt *_router) getGroupPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	http.ServeFile(w, r,
		filepath.Join("/tmp/media", ps.ByName("chat_id"), "1.jpg"))
}
