package main

import (
	"tags/db"
	"tags/server"
)

func main() {
	db := db.New()
	s := server.Create(db)
	s.Start()
}
