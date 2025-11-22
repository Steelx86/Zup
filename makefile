all: 
	go build -o zup ./cmd/main

run: 
	go run ./cmd/main

test:
	go test ./tests/...

clean:
	rm -rf *.zup zup