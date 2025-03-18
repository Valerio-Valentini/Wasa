package api

import (
	"bytes"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"strings"
	"path/filepath" 
)

func (rt *_router) putNewUserPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	aut := r.Header.Get("Authorization")
	valori := strings.Split(aut, " ")
	user_id := valori[1]

	w.Header().Set("Content-Type", "application/json")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// ctx.Logger.WithError(err).Error("Can't retrieve photo data")
		return
	}
	r.Body = io.NopCloser(bytes.NewBuffer(data))
	path := filepath.Join("/tmp/media", user_id, "1.jpg")
	_, err = os.Stat(path)
	if err == nil {
		os.Remove(path)
	}
	// creare file
	out, err := os.Create(path)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// ctx.Logger.WithError(err).Error("Can't retrieve photo data")
		return
	}
	_, err = io.Copy(out, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// ctx.Logger.WithError(err).Error("Can't retrieve photo data")
		return
	}
	out.Close()
	w.WriteHeader(http.StatusOK) // risposta

}
