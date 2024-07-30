.PHONY: run
run:
	go run cmd/main.go

.PHONY: dev
dev:
	npx nodemon --exec go run ./cmd/main.go --watch .  --signal SIGTERM

.PHONY: test
test:
	go test -v ./...

.PHONY: test-coverage
test-coverage:
	go test -cover ./...