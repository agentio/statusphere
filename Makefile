all:
	go install ./...


lexgen:
	lexgenlite --build-file api.json lexicons
