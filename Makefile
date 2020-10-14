all:

build:
	go build -v -o ./bin/gohan ./cmd/gohan/...

run:
	go run ./cmd/gohan/...
