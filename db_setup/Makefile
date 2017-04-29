ADMIN_PASSWORD=tesselation
HEXGAME_DATABASE=hexgame
HEXGAME_USERNAME=hexgame
HEXGAME_PASSWORD=ragnar
MONGO_ADDRESS=127.0.0.1
MONGO_LOAD=mongo --username=${HEXGAME_USERNAME} --password=${HEXGAME_PASSWORD} ${MONGO_ADDRESS}/${HEXGAME_DATABASE}

mongo-service:
	docker run -d --network host --name hexgame-db mongo --auth ;

mongo-admin-password:
	docker exec -it hexgame-db mongo admin --eval "db.createUser({user: 'admin',pwd: '${ADMIN_PASSWORD}', roles: [ { role: 'userAdminAnyDatabase', db: 'admin' } ] } );"

mongo-create-hexgame-db:
	mongo --username admin --password=${ADMIN_PASSWORD} ${MONGO_ADDRESS}/admin init_db_user.js
	mongo --username hexgame --password=${HEXGAME_PASSWORD} ${MONGO_ADDRESS}/hexgame init_db_collections.js


db-ip-address:
	docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' hexgame-db

clean-db:
	docker stop hexgame-db ; docker rm hexgame-db

load-games:
	${MONGO_LOAD} ../data/clear/game.json
	${MONGO_LOAD} ../data/clear/map.json
	${MONGO_LOAD} ../data/warpwar/game.json
	${MONGO_LOAD} ../data/warpwar/map.json
	${MONGO_LOAD} ../data/blackhole/game.json
	${MONGO_LOAD} ../data/blackhole/map.json
#	mongo --username hexgame --password=ragnar ${MONGO_ADDRESS}/hexgame ../data/games.json
