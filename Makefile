include .env

.env:
	cp .env.sample .env

.PHONY: up down

up:
	docker compose up -d app-api app-worker kafka-ui
	docker compose run --rm migrate -path /migrations/ -database ${DB_URL} up

down:
	docker compose down --remove-orphans --rmi all