# Useful command while working with the DB, requires mongo installed on the host

F3API_MONGO_IP := 127.0.0.1
F3API_MONGO_PORT := 27017
F3API_MONGO_HOST := ${F3API_MONGO_IP}:${F3API_MONGO_PORT}
F3API_MONGO_USER := root
F3API_MONGO_PASS := demo
F3API_MONGO_WAIT := 20

MONGO_AUTH = -u ${F3API_MONGO_USER} -p ${F3API_MONGO_PASS} --authenticationDatabase admin

mongo-login:
	@${COMPOSE} exec db mongo ${MONGO_AUTH}

mongo-status:
	@${COMPOSE} exec db mongo ${MONGO_AUTH} --eval "printjson(db.serverStatus())"

mongo-list-dbs:
	@${COMPOSE} exec db mongo ${MONGO_AUTH} --eval "db.adminCommand({listDatabases: 1})"

mongo-payments:
	@${COMPOSE} exec db mongo mongo mongodb://127.0.0.1:27017/f3api ${MONGO_AUTH} --eval "db.payments.find().pretty()"

mongo-wait-ready:
	@ops/scripts/wait-for-db.sh ${F3API_MONGO_IP} ${F3API_MONGO_PORT} ${F3API_MONGO_WAIT}