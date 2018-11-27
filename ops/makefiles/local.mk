.PHONY: all

COMPOSE_FILE := ops/docker-compose.yml
COMPOSE_PROJECT_NAME := "payments"
COMPOSE = docker-compose -p ${COMPOSE_PROJECT_NAME} -f ${COMPOSE_FILE}

dev:
	@${COMPOSE} up -d

stop:
	@${COMPOSE} stop

down:
	@${COMPOSE} down

rm:
	@docker volume rm payments_storage

logs:
	@${COMPOSE} logs -f --tail="10"

# Example: make enter service=db
enter:
	@${COMPOSE} exec ${service} /bin/bash