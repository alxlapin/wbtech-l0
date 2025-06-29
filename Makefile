ifeq ("$(wildcard .env)","")
$(warning .env file not found, creating from sample)
ENV_FILE := $(shell cp .env.sample .env && echo .env)
else
ENV_FILE := .env
endif

include $(ENV_FILE)

.PHONY: up down init-env

up:
	docker compose up -d app postgres
	docker compose run --rm migrate -path /migrations/ -database ${DB_URL} up

down:
	docker compose down --remove-orphans --rmi