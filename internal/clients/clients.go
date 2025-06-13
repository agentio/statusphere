package clients

import (
	"log"
	"os"

	xrpc_anonymous "github.com/agentio/atiquette/pkg/xrpc/anonymous"
	xrpc_session "github.com/agentio/atiquette/pkg/xrpc/session"
)

// This should point to the ATProto calling port of your IO proxy.

func proxy() string {
	p := os.Getenv("ATPROTO_PROXY")
	if p != "" {
		return p
	}
	return "http://localhost:7000"
}

var SessionClient = &xrpc_session.Client{Proxy: proxy()}
var AnonymousClient = &xrpc_anonymous.Client{Host: "https://public.api.bsky.app"}

func init() {
	log.Printf("running with proxy %s", proxy())
}
