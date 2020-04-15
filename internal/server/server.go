package server

import (
	"encoding/json"
	"fmt"
	"gojson/internal/common"
	"log"
	"net/http"
)

func generalHandle(val interface{}) http.HandlerFunc {
	res, err := json.Marshal(val)
	common.Check(err)
	return func(w http.ResponseWriter, r *http.Request) {
		w.Write(res)
	}
}

// Parse generate handle URL's and seperate data
func Parse(all map[string]interface{}) {
	for path, val := range all {
		fmt.Println(path)
		http.HandleFunc(fmt.Sprintf("/%s", path), generalHandle(val))
	}
}

// Serve with port and host address
func Serve(host string, port string) {
	err := http.ListenAndServe(fmt.Sprintf("%s:%s", host, port), logRequest(http.DefaultServeMux))
	common.Check(err)
}

func logRequest(handler http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s\n", r.Method, r.URL)
		handler.ServeHTTP(w, r)
	})
}
