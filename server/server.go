package server

import (
	"fmt"
	"log"
	"net/http"
	"tags/db"
	"tags/server/routes"

	"github.com/julienschmidt/httprouter"
)

type Server struct {
	port   int
	router *httprouter.Router
}

func Create(db *db.DB) *Server {
	rs := routes.New(db)

	router := httprouter.New()

	router.POST("/count", rs.CountTags)
	router.POST("/create", rs.CreateTag)

	port := 8080
	return &Server{port, router}
}

func (s *Server) Start() {
	log.Printf("server running on port %d", s.port)
	serverPort := fmt.Sprintf(":%d", s.port)
	log.Fatal(http.ListenAndServe(serverPort, s.router))
}
