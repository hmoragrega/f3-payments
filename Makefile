.PHONY: all

COMPOSE_FILE = ops/docker-compose.yml
COMPOSE_PROJECT_NAME = payments
COMPOSE_COMMAND = docker-compose -p ${COMPOSE_PROJECT_NAME} -f ${COMPOSE_FILE}

dev:
	@${COMPOSE_COMMAND} up -d

stop:
	@${COMPOSE_COMMAND} stop

down:
	@${COMPOSE_COMMAND} down

logs:
	@${COMPOSE_COMMAND} logs -f --tail="10"

# Example: make enter service=db
enter:
	@@${COMPOSE_COMMAND} exec ${service} bash

include ops/makefiles/*.mk