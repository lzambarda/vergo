.PHONY: lint
lint:
	@go fmt ./...
	@go vet ./...
	@golangci-lint run ./...

.PHONY: dependencies
dependencies: ## Install dependencies needed to work with this repo
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.0
	@go mod tidy
