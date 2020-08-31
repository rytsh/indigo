package server

import (
	"compress/gzip"
	"fmt"
	"gojson/internal/common"
	"gojson/internal/reader"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
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
		// trim slash chars
		r.URL.Path = "/" + strings.Trim(r.URL.Path, "/")
		searchURL := strings.Trim(strings.TrimPrefix(r.URL.Path, "/"+common.API), "/")
		// parse
		var searchURLX []string
		if searchURL != "" {
			searchURLX = strings.Split(searchURL, "/")
		}
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
			if r.URL.Path == "/"+common.API {
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
func SetHandle() error {
	// Serve API
	if common.NoAPI == false {
		reg, _ := regexp.Compile(fmt.Sprintf("(?i)/%s.*(([/]+.*)*)$", common.API))
		mux.HandleFunc(reg, generalHandle(&reader.All))
	}

	// Serve Static Folder
	if common.StaticFolder != "" {
		reg, _ := regexp.Compile("(?i)/.*$")
		fs := http.FileServer(http.Dir(common.StaticFolder))
		mux.Handler(reg, fs)
	}

	if common.NoAPI && common.StaticFolder == "" {
		return fmt.Errorf("Nothing to serve")
	}
	return nil
}

// Gzip Compression
type gzipResponseWriter struct {
	io.Writer
	http.ResponseWriter
}

func (w gzipResponseWriter) Write(b []byte) (int, error) {
	return w.Writer.Write(b)
}

func logRequest(handler http.Handler) http.Handler {
	authUserPass := strings.Split(common.AuthBasic, ":")
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Basic Auth
		if len(authUserPass) > 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

			username, password, authOK := r.BasicAuth()
			if authOK == false {
				http.Error(w, "Not authorized", 401)
				return
			}

			if username != authUserPass[0] || password != authUserPass[1] {
				http.Error(w, "Not authorized", 401)
				return
			}
		}
		// End Basic Auth

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			handler.ServeHTTP(w, r)
		} else {
			w.Header().Set("Content-Encoding", "gzip")
			gz := gzip.NewWriter(w)
			defer gz.Close()
			gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
			handler.ServeHTTP(gzw, r)
		}
		log.Printf("- %s - %s %s\n", r.RemoteAddr[:strings.Index(r.RemoteAddr, ":")], r.Method, r.URL)
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
