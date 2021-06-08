OUTPUT=bin/findme
all: clean build test

build:
	go build -o $(OUTPUT) -v

clean:
	@rm -f $(OUTPUT) coverage.*

run: build
	$(OUTPUT)

test:
	go test ./...

dev: clean build test run

coverage: 
	go test -coverprofile=coverage.out -cover ./... 
	go tool cover -html=coverage.out -o coverage.html
	xdg-open coverage.html
