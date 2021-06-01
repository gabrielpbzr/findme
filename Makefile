OUTPUT=bin/findme
all: clean build test

build:
	go build -o $(OUTPUT) -v

clean:
	@rm -f $(OUTPUT)

run: build
	$(OUTPUT)

test:
	go test ./...

dev: clean build test run