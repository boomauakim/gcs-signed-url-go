run:
	go run cmd/main.go
dev:
	npx nodemon --exec go run ./cmd/main.go --watch .  --signal SIGTERM