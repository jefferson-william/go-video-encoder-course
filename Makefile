install:
	go mod tidy

run:
	go run main.go

test:
	go test ./test/...

unit:
	go test ./test/unit

memory:
	go test ./test/memory