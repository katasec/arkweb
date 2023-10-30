package handlers

import (
	"embed"
	"io/fs"
	"net/http"
	"strings"
)

//go:embed all:assets
var assets embed.FS

func FileHandler(http.HandlerFunc) http.HandlerFunc {

	var prefix = "/"
	// Use the file system to serve static files
	assets := getAssets()
	fs := http.FileServer(http.FS(assets))

	return func(w http.ResponseWriter, r *http.Request) {

		strippedRequest := r.Clone(r.Context())
		strippedRequest.URL.Path = strings.TrimPrefix(strippedRequest.URL.Path, prefix)

		fs.ServeHTTP(w, r)
	}
}

func getAssets() fs.FS {
	fs, err := fs.Sub(assets, "assets")
	if err != nil {
		panic(err)
	}

	return fs
}
