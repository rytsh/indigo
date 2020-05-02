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

func generalHandle(val *interface{}) http.HandlerFunc {
	// TODO: index id

	return func(w http.ResponseWriter, r *http.Request) {
		// set return type
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// trim end of dangling slash chars
		r.URL.Path = common.TrimSuffixAll(r.URL.Path, '/')
		// get inner url
		rVal, rAVal, rAIndex, err := reader.GoInner(val, strings.Split(r.URL.Path, "/")[1:])
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, err.Error())
			return
		}

		successMsg := `{"msg":"success"}`
		switch r.Method {
		case http.MethodGet:
			if r.URL.Path == "" {
				// Welcome home
				// TODO: add an UI webcontente
				w.Header().Set("Content-Type", "text/plain; charset=utf-8")
				fmt.Fprintf(w, common.GetInfo())
			} else {
				result, err := reader.GetHandle(rVal)
				if err != nil {
					w.WriteHeader(http.StatusBadRequest)
					fmt.Fprintf(w, err.Error())
				} else {
					w.WriteHeader(http.StatusOK)
					w.Write(result)
				}
			}
		case http.MethodPost:
			err := reader.PostHandle(rVal, r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, err.Error())
			} else {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, successMsg)
			}
		case http.MethodPut:
			err := reader.PutHandle(rVal, r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, err.Error())
			} else {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, successMsg)
			}
		case http.MethodPatch:
			err := reader.PatchHandle(rVal, r.Body)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, err.Error())
			} else {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, successMsg)
			}
		case http.MethodDelete:
			err := reader.DeleteHandle(rVal, rAVal, rAIndex)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, err.Error())
			} else {
				w.WriteHeader(http.StatusOK)
				fmt.Fprintf(w, successMsg)
			}
		default:
			w.WriteHeader(http.StatusBadRequest)
			fmt.Fprintf(w, `{"err": "GET, POST, PUT, PATCH, DELETE supported!"`)
		}
	}
}

// SetHandle generate handle URL's
func SetHandle() {
	reg, _ := regexp.Compile("(?i)/.*(([/]+.*)*)$")
	mux.HandleFunc(reg, generalHandle(&reader.All))
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
