build:
	@go build -o ./bin/goPoker

run: build
	@./bin/goPoker