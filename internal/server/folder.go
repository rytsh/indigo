package server

import (
	"indigo/internal/common"
	"net/http"
	"strings"
)

// FolderHandle handle function of FOLDER
func FolderHandle(folder http.Dir, path *string) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.URL.Path == common.FolderPath && common.FolderPath != "/" {
			LocalRedirect(w, r, r.URL.Path+"/")
			return
		}
		if *path != "/" {
			r.URL.Path = strings.TrimPrefix(r.URL.Path, *path)
		}
		FileServer(folder).ServeHTTP(w, r)
	}
}
