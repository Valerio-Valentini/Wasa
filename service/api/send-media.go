package api

import (
	"github.com/julienschmidt/httprouter"
	"net/http"
	"image/jpeg"
	"image/png"
)

func (rt *_router) sendMedia(w http.ResponseWriter, r *http.Request, ps httprouter.Params, ctx reqcontext.RequestContext)
{
	aut := r.Header.Get("Authorization")
	valori := strings.Split(aut, " ")
	user_id := valori[0]

	w.Header().Set("Content-Type", "application/json")
	data, err := io.ReadAll(r.Body)
	if err != nil
		{
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("Can't retrieve photo data")
			return
		}
	r.Body = io.NopCloser(bytes.NewBuffer(data))
	id, err := rt.db.GetIdPhoto(user_id)
	if err != nil
		{
			w.WriteHeader(http.StatusInternalServerError)
			ctx.Logger.WithError(err).Error("session: can't create response json")
			return
		}
	
	//creare file
	out, err := os.Create("./media/" + id)
	if err != nil
	{
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Can't retrieve photo data")
		return
	}
	_,err = io.Copy(out, r.Body)
	if err != nil
	{
		w.WriteHeader(http.StatusInternalServerError)
		ctx.Logger.WithError(err).Error("Can't retrieve photo data")
		return
	}
	out.Close()
	w.WriteHeader(http.StatusOk)
}