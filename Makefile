.PHONY:install
install:
	@go install github.com/golangci/golangci-lint/cmd/golangci-lint@v1.44.2

.PHONY:lint
lint:
	@golangci-lint -v run ./...

.PHONY:test
test:
	@go clean -testcache
	@go test -v ./... | grep -v "_mock.go"

.PHONY:coverage
coverage:
	@go clean -testcache
	@go test -coverprofile=tmp/coverage.out ./...
	@go tool cover -html=tmp/coverage.out

.PHONY:run
run:
	@docker-compose up --build