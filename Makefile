build:
	go build -o bin/taulang .

clean:
	rm -rf ./bin

test:
	go test -v ./...

test-coverage:
	go test -cover ./...

test-coverage-html:
	go test -coverprofile=coverage.out ./...
	go tool cover -html=coverage.out -o coverage.html