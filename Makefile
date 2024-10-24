# rundb: run only the db container (used for debugging in local run)
.PHONY: rundb
rundb:
	docker compose up -d --build cc-db

# rebuild: rebuild all containers
.PHONY: rebuild
rebuild:
	docker compose up -d --build

# rebuild: rebuild frontend container
.PHONY: rebuild-front
rebuild-front:
	docker compose up -d --build cc-frontend

## startd: start docker container for server
.PHONY: compose
compose:
	docker compose up -d 