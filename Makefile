all:
	go install ./...


lexgen:
	lexgenlite --build-file build/xyz.json lexicons
