# LOCAL

install:
	go mod download

run:
	go run ./cmd/dota2/main.go

dev:
	gin -p 3001 -a 3006 -i -d cmd/dota2/ run main.go