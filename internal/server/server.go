package server

import (
	"compress/gzip"
	"fmt"
	"indigo/internal/common"
	"indigo/internal/reader"
	"io"
	"log"
	"net/http"
	"strings"
)

// SRV is a general http server
var SRV http.Server

var mux = &MxHandler{}

var idleConnsClosed = make(chan struct{})

// setHandle generate handle URL's
func setHandle() error {
	// UI
	if common.NoUI == false {
		SetRegexString(common.UIPath, "UI")
		mux.HandleFunc(&checkAll.UI, UIHandle())
	}

	// Serve API
	if common.NoAPI == false {
		SetRegexString(common.APIPath, "API")
		mux.HandleFunc(&checkAll.API, APIHandle(&reader.All))
	}

	// Serve Static Folder
	if common.StaticFolder != "" {
		SetRegexString(common.FolderPath, "FOLDER")
		mux.HandleFunc(&checkAll.FOLDER, FolderHandle(http.Dir(common.StaticFolder), &common.FolderPath))
	}

	if common.NoUI && common.NoAPI && common.StaticFolder == "" {
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

func preOperation(handler http.Handler) http.Handler {
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

		if strings.Contains(r.Header.Get("Accept-Encoding"), "gzip") {
			w.Header().Set("Content-Encoding", "gzip")
			gz := gzip.NewWriter(w)
			defer gz.Close()
			gzw := gzipResponseWriter{Writer: gz, ResponseWriter: w}
			handler.ServeHTTP(gzw, r)
		} else {
			handler.ServeHTTP(w, r)
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
	SRV.Handler = preOperation(mux)
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
