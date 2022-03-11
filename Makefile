GO_BIN := go

.PHONY: dev-golib-on
dev-golib-on:
	@go mod edit -replace github.com/lonepeon/golib=../golib
	@go mod download
	@go mod vendor

.PHONY: dev-golib-off
dev-golib-off:
	@go mod edit -dropreplace github.com/lonepeon/golib
	@go mod download
	@go mod vendor

.PHONY: test-generate
test-generate:
	@echo $@
	@./scripts/assert-generated-files-updated.sh

.PHONY: test
test: test-unit test-integration test-format test-lint test-security

.PHONY: test-acceptance
test-acceptance:
	@echo $@

.PHONY: test-acceptance-deps
test-acceptance-deps:
	@echo $@

.PHONY: test-integration
test-integration:
	@echo $@
	@$(GO_BIN) test ./... -run ^TestIntegration

.PHONY: test-lint
test-lint:
	@echo $@
	@$(GO_BIN) run ./vendor/github.com/golangci/golangci-lint/cmd/golangci-lint run

.PHONY: test-format
test-format:
	@echo $@
	@data=$$(gofmt -l main.go internal);\
		 if [ -n "$${data}" ]; then \
			>&2 echo "format is broken:"; \
			>&2 echo "$${data}"; \
			exit 1; \
		 fi

.PHONY: test-security
test-security:
	@echo $@
	@$(GO_BIN) run ./vendor/honnef.co/go/tools/cmd/staticcheck/staticcheck.go

.PHONY: test-unit
test-unit:
	@echo $@
	@$(GO_BIN) test -short ./...
