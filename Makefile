.PHONY:
.SILENT:

build:
	go build ./cmd/api/main.go

run:
	docker-compose up --build

test:
	go test -v ./...
