all:

build:
	go build -v -o ./bin/gohan ./cmd/gohan/...

run:
	go run ./cmd/gohan/...

compose:
	docker-compose up -d --build

exec-sources:
	docker exec -it gohan.sources bash

# show db migrations status
db-status:
	migrate -path migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose version

# apply up migration
db-up:
	migrate -path migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose up

# apply down migration
db-down:
	migrate -path migrations -database "postgresql://$(DB_USER):$(DB_PASSWORD)@$(DB_HOST):$(DB_PORT)/$(DB_NAME)?sslmode=disable" -verbose down

# generate golang models based on queries from ./internal/db/query
db-generate:
	sqlc generate
