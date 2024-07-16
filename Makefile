.PHONY: build
build: # Build image
	@docker-compose build

.PHONY: up
up: # Up docker-compose
	@docker-compose up -d
	@docker-compose logs app -f --no-log-prefix

.PHONY: down
down: # Down docker-compose
	@docker-compose down

.PHONY: stop
stop: # Stop docker-compose
	@docker-compose stop

.PHONY: run
run: # Run app in local
	@CONFIG=local go run ./cmd/app/main.go
