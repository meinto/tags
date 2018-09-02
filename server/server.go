package server

type Server struct{}

func Create() *Server {
	// initialize db
	// initialize routes
	return &Server{}
}

func (s *Server) Start() {}
