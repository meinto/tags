package routes

import (
	"net/http"
	"tags/db"

	"github.com/julienschmidt/httprouter"
)

type Routes struct {
	db *db.DB
}

func New(db *db.DB) *Routes {
	return &Routes{db}
}

func (rs *Routes) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("tags - a go app to handle tags generically"))
}
