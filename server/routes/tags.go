package routes

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func (rs *Routes) CountTags(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {

}

func (rs *Routes) CreateTag(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	body := r.Body
	defer body.Close()

	bodyBytes, err := ioutil.ReadAll(body)
	if err != nil {
		w.Header().Add("status", fmt.Sprintf("%d", http.StatusInternalServerError))
		w.Write([]byte("error occured while reading body :("))
	}
	tag := rs.db.CreateTag(string(bodyBytes))

	enc := json.NewEncoder(w)
	enc.Encode(tag)
}
