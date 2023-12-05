.PHONY: run
run: build
	./${{ application_name }}

.PHONY: build
build:
	go build -o ${{ application_name }} .

.PHONY: test
test:
	go test -v ./...

.PHONY: lint
lint:
	golangci-lint run
