.PHONY:
.SILENT:

build:
	go build ./cmd/api/main.go

run: build
	docker-compose up --build

test:
	go test -v ./...
