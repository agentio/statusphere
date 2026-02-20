all:	xrpc statusphere

statusphere:
	go install ./...

xrpc:
	slink generate xrpc -i lexicons -i xyz -m xrpc.json
