package server

import (
	"net/http"
	"strings"
)

type route struct {
	pattern *checkSelection
	handler http.Handler
}

// MxHandler is a main mux
type MxHandler struct {
	routes []*route
}

// Handler add routes with handler
func (h *MxHandler) Handler(pattern *checkSelection, handler http.Handler) {
	h.routes = append(h.routes, &route{pattern, handler})
}

// HandleFunc add routes with func
func (h *MxHandler) HandleFunc(pattern *checkSelection, handler func(http.ResponseWriter, *http.Request)) {
	h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handler)})
}

func (h *MxHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.useCheck == regexSelect {
			if route.pattern.regexCheck.MatchString(r.URL.Path) {
				route.handler.ServeHTTP(w, r)
				return
			}
		} else {
			if strings.HasPrefix(r.URL.Path, route.pattern.stringCheck) {
				route.handler.ServeHTTP(w, r)
				return
			}
		}
	}
	// no pattern matched; send 404 response
	http.Error(w, "Not existed path!", http.StatusNotFound)
}
