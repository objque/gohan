all:

build:
	go build -v -o ./bin/gohan ./cmd/gohan/...

run:
	go run ./cmd/gohan/...

compose:
	docker-compose up -d --build

exec-sources:
	docker exec -it gohan.sources bash
