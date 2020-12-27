package server

import (
	"fmt"
	"indigo/internal/common"
	"net/http"
)

// UIHandle serve UI content
func UIHandle() http.HandlerFunc {
	// TODO: index id
	return func(w http.ResponseWriter, r *http.Request) {
		// Welcome home
		// TODO: add an UI webcontente
		r.URL.Path = common.TrimSlash(r.URL.Path)
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, common.GetInfo())
	}
}
