# Useful command while working with the DB, requires mongo installed on the host

F3_API_MONGO_USER := root
F3_API_MONGO_PASS := demo
F3_API_MONGO_DB := f3api

MONGO_WAIT = 20
MONGO_ADMIN = ${COMPOSE_COMMAND} exec db mongo mongodb://127.0.0.1:27017/admin -u ${F3_API_MONGO_USER} -p ${F3_API_MONGO_PASS}
MONGO_F3API = ${COMPOSE_COMMAND} exec db mongo mongodb://127.0.0.1:27017/${F3_API_MONGO_DB} --authenticationDatabase admin
FIXTURES = payments.json

mongo-test:
	@${MONGO_COMMAND} 

mongo-login:
	@${MONGO_ADMIN}

mongo-status:
	@${MONGO_ADMIN} --eval "printjson(db.serverStatus())"

mongo-list-dbs:
	@${MONGO_ADMIN} --eval "db.adminCommand({listDatabases: 1})"

mongo-payments:
	@${MONGO_F3API} --eval "db.payments.find().pretty()"

mongo-load-fixtures:
	# Circle remote docker setup don't allow mount volumes, so we copy the fixtures
	@docker cp ops/mongo/${FIXTURES} payments_db_1:/${FIXTURES}  
	@${COMPOSE_COMMAND} exec db mongoimport -u ${F3_API_MONGO_USER} -p ${F3_API_MONGO_PASS} --authenticationDatabase admin --db ${F3_API_MONGO_DB} --collection payments --drop --jsonArray --file /${FIXTURES} 