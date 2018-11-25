# Useful command while working with the DB, requires mongo installed on the host

F3API_MONGO_IP := 127.0.0.1
F3API_MONGO_PORT := 27017
F3API_MONGO_HOST := ${F3API_MONGO_IP}:${F3API_MONGO_PORT}
F3API_MONGO_USER := root
F3API_MONGO_PASS := demo
F3API_MONGO_WAIT := 20

mongo-login:
	@mongo mongodb://${F3API_MONGO_HOST}/admin -u ${F3API_MONGO_USER} -p '${F3API_MONGO_PASS}'

mongo-status:
	@mongo mongodb://${F3API_MONGO_HOST}/admin -u ${F3API_MONGO_USER} -p '${F3API_MONGO_PASS}' --eval "printjson(db.serverStatus())"

mongo-list-dbs:
	@mongo mongodb://${F3API_MONGO_HOST}/admin -u ${F3API_MONGO_USER} -p '${F3API_MONGO_PASS}' --eval "db.adminCommand({listDatabases: 1})"

mongo-payments:
	@mongo mongodb://${F3API_MONGO_HOST}/f3api -u ${F3API_MONGO_USER} -p '${F3API_MONGO_PASS}' --authenticationDatabase admin --eval "db.payments.find().pretty()"

mongo-wait-ready:
	@ops/scripts/wait-for-db.sh ${F3API_MONGO_IP} ${F3API_MONGO_PORT} ${F3API_MONGO_WAIT}

mongo-drop-db:
	@mongo mongodb://${F3API_MONGO_HOST}/f3api -u ${F3API_MONGO_USER} -p '${F3API_MONGO_PASS}' --authenticationDatabase admin --eval "db.payments.drop()"

mongo-load-fixtures: mongo-drop-db
	@mongoimport --host ${F3API_MONGO_HOST} -u ${F3API_MONGO_USER} -p ${F3API_MONGO_PASS} --authenticationDatabase admin --db f3api --collection payments --jsonArray --file ops/mongo/payments.json