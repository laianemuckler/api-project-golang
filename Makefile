build:
	@go build -o bin/api-project-golang

run: build
		 @./bin/api-project-golang

#Run compose
compose-up:
	@docker-compose -f docker-compose.yml --env-file .env up -d

compose-down:
	@docker-compose -f docker-compose.yml --env-file .env down
