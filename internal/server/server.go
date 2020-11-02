package server

import (
	"compress/gzip"
	"fmt"
	"indigo/internal/common"
	"indigo/internal/reader"
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

// setHandle generate handle URL's
func setHandle() error {
	// Set regex strings
	SetRegexs()

	// UI
	if common.NoUI == false {
		reg := regexp.MustCompile(regex.UI)
		mux.HandleFunc(reg, UIHandle(), false, nil)
	}

	// Serve API
	if common.NoAPI == false {
		reg := regexp.MustCompile(regex.API)
		mux.HandleFunc(reg, APIHandle(&reader.All), false, nil)
	}

	// Serve Static Folder
	if common.StaticFolder != "" {
		reg := regexp.MustCompile(regex.FOLDER)
		fs := http.FileServer(http.Dir(common.StaticFolder))
		mux.Handler(reg, fs, true, &common.FolderPath)
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
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Basic Auth
		if len(common.AuthBasic) > 1 {
			w.Header().Set("WWW-Authenticate", `Basic realm="Restricted"`)

			username, password, authOK := r.BasicAuth()

			if authOK == false || username != common.AuthBasic[0] || password != common.AuthBasic[1] {
				http.Error(w, "Not authorized", 401)
				return
			}
		}
		// End Basic Auth

		// Clean URL
		r.URL.Path = common.TrimSlash(r.URL.Path)

		if !strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			handler.ServeHTTP(w, r)
		} else {
			w.Header().Set("Content-Encoding", "gzip")
			gz := gzip.NewWriter(w)
			defer gz.Close()
			gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
			handler.ServeHTTP(gzw, r)
		}
		log.Printf("- %s - %s %s\n", r.RemoteAddr[:strings.LastIndex(r.RemoteAddr, ":")], r.Method, r.URL)
	})
}

// Serve with port and host address
func Serve(host string, port string) {
	if err := setHandle(); err != nil {
		common.ErrorPrintExit(err.Error(), 5)
	}

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
