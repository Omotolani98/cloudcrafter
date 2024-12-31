.PHONY: build release

build:
	@echo "Building CloudCrafter CLI..."
	go build -o cloudcrafter ./cmd/cli

release:
	@echo "Releasing CloudCrafter CLI..."
	goreleaser release --clean