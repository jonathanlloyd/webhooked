all: build

test:
	@echo Running unit tests...
	@go test ./...

check: test
