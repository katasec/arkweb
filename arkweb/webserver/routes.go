package webserver

import (
	"github.com/writeameer/arkweb/webserver/handlers"
)

func (s *Server) initialiseRoutes() {

	// Define the sequence of middleware used to process all requests
	middlewareSequence := MultipleMiddleware(
		handlers.EntryHandler,
		handlers.LogHandler,
		handlers.FileHandler,
	)

	// Use the file system to serve static files
	s.mux.Handle("/", middlewareSequence)

}
