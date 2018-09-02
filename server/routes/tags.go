package routes

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rs *Routes) CreateTag(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body := r.Body
	defer body.Close()

	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occured while reading body :("))
		return
	}
	tag := rs.db.CreateTag(string(bodyBytes))

	enc := json.NewEncoder(w)
	enc.Encode(tag)
}

func (rs *Routes) CountTags(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body := r.Body
	defer body.Close()

	var jsn map[string]interface{}
	if err := json.NewDecoder(body).Decode(&jsn); err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("error occured while decoding body :("))
		return
	}
	nmbr := rs.db.CountTags(jsn)

	enc := json.NewEncoder(w)
	enc.Encode(map[string]int{"count": nmbr})
}
