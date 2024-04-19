DOCKER_COMPOSE_FILE := docker-compose.yml

.PHONY: docker-build-backend docker-run-backend docker-build docker-run

docker-build-backend:
	docker-compose -f $(DOCKER_COMPOSE_FILE) build backend

docker-run-backend:
	docker-compose -f $(DOCKER_COMPOSE_FILE) up

docker-build: docker-build-backend

docker-run: docker-build docker-run-backend