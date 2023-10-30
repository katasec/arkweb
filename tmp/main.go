package main

import (
	"github.com/writeameer/arkweb/webserver"
)

func main() {
	server := webserver.NewServer()
	server.Start()
}
