.PHONY: all

COMPOSE_FILE := ops/docker-compose.yml
COMPOSE_PROJECT_NAME := "payments"

dev:
	@docker-compose -p ${COMPOSE_PROJECT_NAME} -f ${COMPOSE_FILE} up -d

stop:
	@docker-compose -p ${COMPOSE_PROJECT_NAME} -f ${COMPOSE_FILE} stop

down:
	@docker-compose -p ${COMPOSE_PROJECT_NAME} -f ${COMPOSE_FILE} down

rm:
	@docker volume rm payments_storage

logs:
	@docker-compose -p ${COMPOSE_PROJECT_NAME} -f ${COMPOSE_FILE} logs -f --tail="10"

# Example: make enter service=db
enter:
	@docker exec -it ${COMPOSE_PROJECT_NAME}_$(service)_1 /bin/bash