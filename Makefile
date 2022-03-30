GO_BIN := go

ifdef SSH_KNOWN_HOSTS
	SSH_KNOWN_HOST_OPTION := -o UserKnownHostsFile=$(SSH_KNOWN_HOSTS)
endif

.PHONY: deploy
deploy:
ifndef VERSION
	$(error VERSION variable is not set)
endif
ifndef SSH_KEY_PATH
	$(error SSH_KEY_PATH variable is not set)
endif
ifndef DEPLOY_SERVER
	$(error DEPLOY_SERVER variable is not set)
endif
	ssh $(SSH_KNOWN_HOST_OPTION) -i $(SSH_KEY_PATH) $(DEPLOY_SERVER) -- mkdir -p /var/lib/food/$(VERSION); \
	scp $(SSH_KNOWN_HOST_OPTION) -i $(SSH_KEY_PATH) food-linux-amd64 $(DEPLOY_SERVER):/var/lib/food/$(VERSION)/food; \
	ssh $(SSH_KNOWN_HOST_OPTION) -i $(SSH_KEY_PATH) $(DEPLOY_SERVER) -- chown -R app:app /var/lib/food/$(VERSION); \
	ssh $(SSH_KNOWN_HOST_OPTION) -i $(SSH_KEY_PATH) $(DEPLOY_SERVER) -- ln -fs /var/lib/food/$(VERSION)/food /var/lib/food/food; \
	ssh $(SSH_KNOWN_HOST_OPTION) -i $(SSH_KEY_PATH) $(DEPLOY_SERVER) -- systemctl restart food; \
	ssh $(SSH_KNOWN_HOST_OPTION) -i $(SSH_KEY_PATH) $(DEPLOY_SERVER) -- bash -c 'ls --sort time | grep '^20' | tail -n +3 | xargs rm -fr'

.PHONY: food-linux-amd64
food-linux-amd64: food-builder
	@mkdir -p target
	docker run --rm -v "$$PWD/target":/usr/local/bin \
		food-builder:latest \
		cp /go/src/food/$@ /usr/local/bin/$@

.PHONY: food-builder
food-builder:
	docker build -t food-builder:latest .

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
	@docker-compose restart acceptance-tests-runner
	@docker-compose exec -T -- acceptance-tests-runner bash -c 'npm install && npm run test'

.PHONY: test-acceptance-deps
test-acceptance-deps:
	@echo $@
	@docker-compose down
	@docker-compose build
	@docker-compose up --scale webapp=1 --scale acceptance-tests-runner=1 --detach

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
