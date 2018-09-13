package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/meinto/tags/db"
	"github.com/meinto/tags/server/routes"
)

type Server struct {
	port   int
	router *httprouter.Router
}

func Create(db *db.DB) *Server {
	rs := routes.New(db)

	router := httprouter.New()

	router.GET("/", rs.Index)

	router.POST("/create/:index", rs.CreateTag)
	router.POST("/count/:index", rs.CountTags)

	port := 8080
	return &Server{port, router}
}

func (s *Server) Start() {
	log.Printf("server running on port %d", s.port)
	serverPort := fmt.Sprintf(":%d", s.port)
	log.Fatal(http.ListenAndServe(serverPort, s.router))
}
