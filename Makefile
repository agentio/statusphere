all:	xrpc statusphere

statusphere:
	go install ./...

xrpc:
	go install github.com/agentio/slink/cmd/slink-generate@latest
	$(shell go env GOPATH)/bin/slink-generate xrpc -i lexicons-bluesky -i lexicons-local -m xrpc.json
