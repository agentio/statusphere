package clients

import (
	xrpc_anonymous "github.com/agentio/atiquette/pkg/xrpc/anonymous"
	xrpc_session "github.com/agentio/atiquette/pkg/xrpc/session"
)

// This should point to the ATProto calling port of your IO proxy.
const proxy = "http://localhost:7000"

// These are effectively const, but Go won't allow that.
var SessionClient = &xrpc_session.Client{Proxy: proxy}
var AnonymousClient = &xrpc_anonymous.Client{Host: "https://public.api.bsky.app"}
