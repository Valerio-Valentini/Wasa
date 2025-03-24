package api

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/julienschmidt/httprouter"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
)

func (rt *_router) sendMedia(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	aut := r.Header.Get("Authorization")
	valori := strings.Split(aut, " ")
	user_id := valori[1]

	/*
		w.Header().Set("Content-Type", "application/json")
		data, err := io.ReadAll(r.Body)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// ctx.Logger.WithError(err).Error("Can't retrieve photo data")
			return
		}
		r.Body = io.NopCloser(bytes.NewBuffer(data))
		id, err := rt.GetIdPhoto(user_id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			// ctx.Logger.WithError(err).Error("session: can't create response json")
			return
		}

		// creare file
		id_string := strconv.Itoa(id)
		out, err := os.Create("./media/" + id_string)
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
		w.WriteHeader(http.StatusOK)
	*/

	w.Header().Set("Content-Type", "application/json")
	data, err := io.ReadAll(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// ctx.Logger.WithError(err).Error("Can't retrieve photo data")
		return
	}
	r.Body = io.NopCloser(bytes.NewBuffer(data))
	path := filepath.Join("/tmp/media", ps.ByName("chat_id"), "tmp.jpg")
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
	id, err := rt.db.SendMedia(ps.ByName("chat_id"), user_id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		// ctx.Logger.WithError(err).Error("Can't retrieve photo data")
		return
	}
	err = os.Rename(path, filepath.Join("/tmp/media", ps.ByName("chat_id"), fmt.Sprintf("%d.jpg", id)))
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	defer out.Close()
	w.WriteHeader(http.StatusCreated) // risposta
	type Resp struct {
		Resp string `json:"photo_id"`
	}
	answ := strconv.FormatInt(id, 10)
	_ = json.NewEncoder(w).Encode(Resp{Resp: answ})
}
