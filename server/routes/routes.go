package routes

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/meinto/tags/db"
)

const VERSION = "1"

type Routes struct {
	db *db.DB
}

func New(db *db.DB) *Routes {
	return &Routes{db}
}

func (rs *Routes) Index(w http.ResponseWriter, r *http.Request, _ httprouter.Params) {
	w.Write([]byte("tags - a go app to handle tags generically - version: " + VERSION))
}
