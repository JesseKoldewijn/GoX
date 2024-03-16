package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/JesseKoldewijn/go-x/handlers"
)

func makeServerFromMux(mux *http.ServeMux) *http.Server {
	// set timeouts so that a slow or malicious client doesn't
	// hold resources forever
	return &http.Server{
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 5 * time.Second,
		IdleTimeout:  120 * time.Second,
		Handler:      mux,
	}
}

func makeHTTPServer() *http.Server {
	mux := &http.ServeMux{}

	routes := map[string]func(http.ResponseWriter, *http.Request){
		"/":        handlers.HandleIndex,
		"/dynamic": handlers.HandleIndexDynamic,
	}

	for path, handler := range routes {
		mux.HandleFunc(path, handler)
	}

	return makeServerFromMux(mux)

}

func main() {
	port := flag.String("port", "3002", "Port to listen to")
	flag.Parse()

	listeningPort := ":" + *port
	log.Println(listeningPort)

	var httpSrv = makeHTTPServer()

	httpSrv.Addr = listeningPort
	fmt.Printf("Starting HTTP server on %s\n", listeningPort)

	err := httpSrv.ListenAndServe()
	if err != nil {
		log.Fatalf("httpSrv.ListenAndServe() failed with %s", err)
	}
}
