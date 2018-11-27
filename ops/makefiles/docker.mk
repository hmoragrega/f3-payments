.PHONY: all

COMPOSE_FILE := ops/docker-compose.yml
COMPOSE_PROJECT_NAME := "payments"

ifeq ($(CI),true)
	COMPOSE = docker-compose -p ${COMPOSE_PROJECT_NAME} -f ${COMPOSE_FILE} -f ops/docker-compose.ci.yml
else 
	COMPOSE = docker-compose -p ${COMPOSE_PROJECT_NAME} -f ${COMPOSE_FILE}
endif

dev:
	@${COMPOSE} up -d

build:
	@${COMPOSE} build

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
	@${COMPOSE} exec ${service} /bin/sh