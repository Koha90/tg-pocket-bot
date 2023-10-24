build:
	@go build -o bin/tg-pocket-bot cmd/main.go 

run: build
	@./bin/tg-pocket-bot
