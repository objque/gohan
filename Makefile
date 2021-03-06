override RELEASE="$(git tag -l --points-at HEAD)"
override COMMIT="$(shell git rev-parse --short HEAD)"
override BUILD_TIME="$(shell date -u '+%Y-%m-%dT%H:%M:%S')"

all:

build:
	go build -v -o ./bin/gohan ./cmd/gohan/...

test t:
	go test -v ./internal/...

lint l:
	bash ./scripts/revive.sh
	bash ./scripts/golangci-lint.sh

run:
	go run ./cmd/gohan/... --config ./gohan.example.yml

compose:
	docker-compose up -d --build

exec-sources:
	docker exec -it gohan.sources bash

image:
	docker build \
		--build-arg RELEASE=${RELEASE} \
		--build-arg COMMIT=${COMMIT} \
		--build-arg BUILD_TIME=${BUILD_TIME} \
		-t "objque/gohan:latest" .
