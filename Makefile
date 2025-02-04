run:
	go run main.go

deps:
	go install github.com/swaggo/swag/cmd/swag@latest && \
	goproxy=direct \
	GOSUMDB=off \
	go mod tidy

test:
	go test ./...  -cover -coverprofile=coverage.out

watch-coverage:
	go tool cover -html=coverage.out

gen-mocks:
	mockery --dir=./internal --all --output=internal/application/usecases/mocks

check-linter:
	 golangci-lint run --path-prefix=./ -v --skip-dirs constant --config=./golangci-lint.yaml --timeout=5m