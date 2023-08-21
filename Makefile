build:
	@go build -o bin/api-project-golang

run: build
		 @./bin/api-project-golang