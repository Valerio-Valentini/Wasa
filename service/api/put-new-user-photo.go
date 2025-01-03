package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"strings"
	"io"
	"bytes"
	"os"
	"strconv"
	"errors"
)

func (rt *_router) putNewUserPhoto(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
	aut := r.Header.Get("Authorization")
	valori := strings.Split(aut, " ")
	user_id := valori[0]

	w.Header().Set("Content-Type", "application/json")
	data, err := io.ReadAll(r.Body)
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			//ctx.Logger.WithError(err).Error("Can't retrieve photo data")
			return
		}
	r.Body = io.NopCloser(bytes.NewBuffer(data))
	id, err := rt.GetIdProfilePicture(user_id)
	if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			//ctx.Logger.WithError(err).Error("session: can't create response json")
			return
		}
	if (id == 0) {
		id, err = rt.CreateNewId(user_id)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			//ctx.Logger.WithError(err).Error("Can't retrieve photo data")
			return
		}
		 _, err := os.Stat("./media/profile_picture/" + user_id)
		 if os.IsNotExists(err) {
			//creare cartella
			err = os.Mkdir("./media/profile_picture/" + user_id, os.ModeDir)
			if err != nil {
				w.WriteHeader(http.StatusInternalServerError)
				//ctx.Logger.WithError(err).Error("Can't retrieve photo data")
				return
			}
		 }	
	}
	
	//creare file
	id_string := strconv.Itoa(id)
	out, err := os.Create("./media/profile_picture/" + user_id + "/" + id_string)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//ctx.Logger.WithError(err).Error("Can't retrieve photo data")
		return
	}
	_,err = io.Copy(out, r.Body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		//ctx.Logger.WithError(err).Error("Can't retrieve photo data")
		return
	}
	out.Close()
	//w.WriteHeader(http.StatusOk) //risposta

}