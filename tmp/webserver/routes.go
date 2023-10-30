package webserver

import (
	"net/http"

	"github.com/writeameer/arkweb/webserver/handlers"
)

func (s *Server) initialiseRoutes() {

	// Use the file system to serve static files
	assets := Assets()
	fs := http.FileServer(http.FS(assets))
	s.mux.Handle("/", http.StripPrefix("/", fs))

	s.mux.HandleFunc("/home", handlers.HomeHandler())
}
