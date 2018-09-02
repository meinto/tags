package routes

import "tags/db"

type Routes struct {
	db *db.DB
}

func New(db *db.DB) *Routes {
	return &Routes{db}
}
