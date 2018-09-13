package main

import (
	"github.com/meinto/tags/db"
	"github.com/meinto/tags/server"
)

func main() {
	db := db.New()
	s := server.Create(db)
	s.Start()
}
