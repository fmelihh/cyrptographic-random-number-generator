build:
	@go build -o bin/cyrptographic-random-number-generator

run: build
	@./bin/cyrptographic-random-number-generator
