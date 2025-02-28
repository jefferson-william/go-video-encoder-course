install:
	go mod tidy

run:
	go run main.go

unit:
	go test ./test/unit