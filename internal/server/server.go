package server

import (
	"fmt"
	"gojson/internal/common"
	"gojson/internal/reader"
	"log"
	"net/http"
	"regexp"
)

// SRV is a general http server
var SRV http.Server

var mux = &RegexpHandler{}

var idleConnsClosed = make(chan struct{})

func generalHandle(val interface{}) http.HandlerFunc {
	// TODO: index id
	// indexData := make(map[string]int)
	// for i, value := range val.([]interface{}) {
	// 	var s string
	// 	switch v := value.(map[string]interface{})["id"].(type) {
	// 	case string:
	// 		s = v
	// 	default:
	// 		s = fmt.Sprintf("%v", v)
	// 	}
	// 	indexData[s] = i
	// }

	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		switch r.Method {
		case http.MethodGet:
			result, err := reader.GetHandle(&val, r.URL.Path)
			if err != nil {
				fmt.Fprintf(w, err.Error())
			} else {
				w.Write(result)
			}

		case http.MethodPost:
			err := reader.PostHandle(&val, r.URL.Path, r.Body)
			if err != nil {
				fmt.Fprintf(w, err.Error())
			} else {
				fmt.Fprintf(w, "{}")
			}
		case http.MethodPut:
			err := reader.PutHandle(&val, r.URL.Path, r.Body)
			if err != nil {
				fmt.Fprintf(w, err.Error())
			} else {
				fmt.Fprintf(w, "{}")
			}
		case http.MethodDelete:
			err := reader.DeleteHandle(&val, r.URL.Path)
			if err != nil {
				fmt.Fprintf(w, err.Error())
			} else {
				fmt.Fprintf(w, "{}")
			}
		default:
			fmt.Fprintf(w, "GET, POST, PUT, DELETE supported!")
		}
	}
}

// SetHandle generate handle URL's and seperate data
func SetHandle(all map[string]interface{}) []string {
	resList := make([]string, 0)
	for path, val := range all {
		resList = append(resList, path)
		reg, _ := regexp.Compile(fmt.Sprintf("(?i)/%s(([/]+.*)*)$", path))
		mux.HandleFunc(reg, generalHandle(val))
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
		log.Printf("%s %s\n", r.Method, r.URL)
		handler.ServeHTTP(w, r)
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
