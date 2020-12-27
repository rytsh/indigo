package server

import (
	"fmt"
	"indigo/internal/common"
	"indigo/internal/reader"
	"net/http"
)

// APIHandle handle function of API
func APIHandle(val *interface{}) http.HandlerFunc {
	// TODO: index id

	return func(w http.ResponseWriter, r *http.Request) {
		// set return type
		w.Header().Set("Content-Type", "application/json; charset=utf-8")
		// trim and split
		var searchURLX []string
		// Clean URL
		searchURLX = common.TrimURL(common.TrimSlash(r.URL.Path))
		// get inner url
		rVal, rAVal, rAIndex, err := reader.GoInner(val, searchURLX)
		if err != nil {
			w.WriteHeader(http.StatusNotFound)
			fmt.Fprintf(w, err.Error())
			return
		}

		successMsg := `{"msg":"success"}`
		switch r.Method {
		case http.MethodGet:
			result, err := reader.GetHandle(rVal)
			if err != nil {
				w.WriteHeader(http.StatusBadRequest)
				fmt.Fprintf(w, err.Error())
			} else {
				w.WriteHeader(http.StatusOK)
				w.Write(result)
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
