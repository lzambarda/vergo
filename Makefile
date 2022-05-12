.PHONY: lint
lint:
	@go fmt ./...
	@go vet ./...
	@golangci-lint run ./...

.PHONY: dependencies
dependencies: ## Install dependencies needed to work with this repo
	@go install github.com/cespare/reflex@latest
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@latest
	@go mod tidy

.PHONY: unit_test
unit_test:
	@-go test -trimpath -failfast ./...


.PHONY: unit_test_watch
unit_test_watch: unit_test
	reflex --shutdown-timeout=500ms --sequential=false -d none -r "(\.go$$)|(go.mod)|(\.sql$$)" make unit_test
