package main

import (
	"fmt"
	"log"
	"net/http"

	xrpc_anonymous "github.com/agentio/statusphere-io/pkg/xrpc/anonymous"
	xrpc_session "github.com/agentio/statusphere-io/pkg/xrpc/session"
	"github.com/gorilla/mux"
)

// This statusphere app listens on this port.
const port = 3000

// This prefix is used by the IO proxy as the root of its OAuth-related handlers.
const prefix = "@"

// This should point to the ATProto calling port of your IO proxy.
const proxy = "http://localhost:7000"

// These are effectively const, but Go won't allow that.
var sessionClient = &xrpc_session.Client{Proxy: proxy}
var anonymousClient = &xrpc_anonymous.Client{Host: "https://public.api.bsky.app"}

func main() {
	go jetstream()

	mux := mux.NewRouter()
	mux.HandleFunc("/", HomeHandler)
	mux.HandleFunc("/status", StatusHandler)
	mux.HandleFunc("/login", LoginHandler)
	mux.HandleFunc("/public/styles.css", StylesHandler)
	httpd := &http.Server{
		Addr:    fmt.Sprintf(":%d", port),
		Handler: mux,
	}
	if err := httpd.ListenAndServe(); err != nil {
		log.Fatalf("%s", err)
	}
}
