package server

import (
	"net/http"
	"strings"
)

// FolderHandle handle function of FOLDER
func FolderHandle(folder http.Dir, path *string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.URL.Path = "/" + strings.TrimPrefix(r.URL.Path, *path)
		FileServer(folder).ServeHTTP(w, r)
	}
}
