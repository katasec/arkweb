package webserver

import (
	"github.com/writeameer/arkweb/webserver/handlers"
)

func (s *Server) initialiseRoutes() {
	s.mux.HandleFunc("/", handlers.HomeHandler())
}
