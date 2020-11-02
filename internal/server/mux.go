package server

import (
	"net/http"
	"regexp"
	"strings"
)

type route struct {
	pattern *regexp.Regexp
	handler http.Handler
	clean   bool
	path    *string
}

// RegexpHandler is a main mux
type RegexpHandler struct {
	routes []*route
}

// Handler add routes with handler
func (h *RegexpHandler) Handler(pattern *regexp.Regexp, handler http.Handler, clean bool, path *string) {
	h.routes = append(h.routes, &route{pattern, handler, clean, path})
}

// HandleFunc add routes with func
func (h *RegexpHandler) HandleFunc(pattern *regexp.Regexp, handler func(http.ResponseWriter, *http.Request), clean bool, path *string) {
	h.routes = append(h.routes, &route{pattern, http.HandlerFunc(handler), clean, path})
}

func (h *RegexpHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	for _, route := range h.routes {
		if route.pattern.MatchString(r.URL.Path) {
			if route.clean {
				r.URL.Path = "/" + strings.Trim(strings.TrimPrefix(r.URL.Path, *route.path), "/")
			}
			route.handler.ServeHTTP(w, r)
			return
		}
	}
	// no pattern matched; send 404 response
	http.Error(w, "Not existed path!", http.StatusNotFound)
}
