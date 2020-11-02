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
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, common.GetInfo())
	}
}
