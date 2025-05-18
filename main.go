package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/agentio/statusphere/internal/handlers"
	"github.com/agentio/statusphere/internal/jetstream"
	"github.com/gorilla/mux"
)

// This statusphere app listens on this port.
const port = 3000

func main() {
	go jetstream.Listen()

	mux := mux.NewRouter()
	mux.HandleFunc("/", handlers.HomeHandler)
	mux.HandleFunc("/status", handlers.StatusHandler)
	mux.HandleFunc("/signin", handlers.SigninHandler)
	mux.HandleFunc("/public/styles.css", handlers.StylesHandler)
	httpd := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	if err := httpd.ListenAndServe(); err != nil {
		log.Fatalf("%s", err)
	}
}
