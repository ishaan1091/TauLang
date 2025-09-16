build:
	go build -o bin/taulang .

build-check:
	go build -v ./...

clean:
	rm -rf ./bin && rm -f coverage.out coverage.html

test:
	go test ./...

test-coverage:
	go test -cover ./...

test-coverage-html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html