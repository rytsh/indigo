package server

import (
	"fmt"
	"gojson/internal/common"
	"gojson/internal/reader"
	"log"
	"net/http"
	"regexp"
	"strings"
	"time"
)

// SRV is a general http server
var SRV http.Server

var mux = &RegexpHandler{}

var idleConnsClosed = make(chan struct{})

func generalHandle(val *interface{}, path string) http.HandlerFunc {
	// TODO: index id
	// reader.All[path] = &val

	return func(w http.ResponseWriter, r *http.Request) {
		// set return type
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// trim end of dangling slash chars
		r.URL.Path = common.TrimSuffixAll(r.URL.Path, '/')
		// get inner url
		rVal, rAVal, rAIndex, err := reader.GoInner(val, strings.Split(r.URL.Path, "/")[2:])
		if err != nil {
			fmt.Fprintf(w, err.Error())
			return
		}

		switch r.Method {
		case http.MethodGet:
			result, err := reader.GetHandle(rVal)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, err.Error())
			} else {
				w.Write(result)
			}
		case http.MethodPost:
			err := reader.PostHandle(rVal, r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, err.Error())
			} else {
				fmt.Fprintf(w, "{}")
			}
		case http.MethodPut:
			err := reader.PutHandle(rVal, r.Body)
			if err != nil {
				fmt.Fprintf(w, err.Error())
			} else {
				fmt.Fprintf(w, "{}")
			}
		case http.MethodDelete:
			err := reader.DeleteHandle(rVal, rAVal, rAIndex)
			if err != nil {
				fmt.Fprintf(w, err.Error())
			} else {
				fmt.Fprintf(w, "{}")
			}
		default:
			fmt.Fprintf(w, `{"err": "GET, POST, PUT, DELETE supported!"`)
		}
	}
}

// SetHandle generate handle URL's and seperate data
func SetHandle() []string {
	resList := make([]string, 0)
	for path, val := range reader.All {
		resList = append(resList, path)
		reg, _ := regexp.Compile(fmt.Sprintf("(?i)/%s(([/]+.*)*)$", path))
		mux.HandleFunc(reg, generalHandle(val, path))
	}
	// set Home
	reg, _ := regexp.Compile("^/$")
	mux.HandleFunc(reg, home())
	return resList
}

func home() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "text/plain; charset=utf-8")
		fmt.Fprintf(w, common.GetInfo())
	}
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		start := time.Now()
		handler.ServeHTTP(w, r)
		log.Printf("%s %s %v us\n", r.Method, r.URL, time.Since(start).Microseconds())
	})
}

// Serve with port and host address
func Serve(host string, port string) {
	SRV.Addr = fmt.Sprintf("%s:%s", host, port)
	SRV.Handler = logRequest(mux)
	if err := SRV.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Fatalf("HTTP server ListenAndServe: %v", err)
	} else {
		log.Println("Server shutdown")
	}
	<-idleConnsClosed
}

// Close server
func Close() {
	close(idleConnsClosed)
}
