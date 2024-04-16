build:
	go build -o ./bin/blockchain-in-go

run: build
	./bin/blockchain-in-go

test:
	go test ./...

testv:
	go test -v ./...