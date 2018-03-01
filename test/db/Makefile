ADMIN_PASSWORD=tesselation
HEXGAME_DATABASE=hexgame
HEXGAME_USERNAME=hexgame
HEXGAME_PASSWORD=ragnar
MONGO_ADDRESS=127.0.0.1
MONGO_LOAD=mongo --quiet --username=${HEXGAME_USERNAME} --password=${HEXGAME_PASSWORD} ${MONGO_ADDRESS}/${HEXGAME_DATABASE}


usage:
	@echo "make mongo-service"
	@echo "make mongo-admin-password"
	@echo "make create-hexgame-db"
	@echo
	@echo "make db-ip-address"
	@echo
	@echo "make load-games"

mongo-service: 
	docker run -d --network host --name hexgame-db mongo --auth
	sleep 3
	touch mongo-service

mongo-admin-password: mongo-service
	docker exec -it hexgame-db mongo admin --eval "db.createUser({user: 'admin',pwd: '${ADMIN_PASSWORD}', roles: [ { role: 'userAdminAnyDatabase', db: 'admin' } ] } );" 
	touch mongo-admin-password

mongo-create-hexgame-db: mongo-admin-password
	mongo --username admin --password=${ADMIN_PASSWORD} ${MONGO_ADDRESS}/admin js/init_db_user.js
	mongo --username hexgame --password=${HEXGAME_PASSWORD} ${MONGO_ADDRESS}/hexgame js/init_db_collections.js
	touch mongo-create-hexgame-db

db-ip-address:
	docker inspect --format='{{range .NetworkSettings.Networks}}{{.IPAddress}}{{end}}' hexgame-db

load-games: mongo-create-hexgame-db
	${MONGO_LOAD} ../data/test/game.json
#	${MONGO_LOAD} ../data/test/map.json
	${MONGO_LOAD} ../data/clear/game.json
#	${MONGO_LOAD} ../data/clear/map.json
	${MONGO_LOAD} ../data/warpwar/game.json
#	${MONGO_LOAD} ../data/warpwar/map.json
	${MONGO_LOAD} ../data/blackhole/game.json
#	${MONGO_LOAD} ../data/blackhole/map.json
	${MONGO_LOAD} ../data/players.json
#	${MONGO_LOAD} ../data/matches.json
#	${MONGO_LOAD} ../data/match_maps.json
#	mongo --username hexgame --password=ragnar ${MONGO_ADDRESS}/hexgame ../data/games.json
	touch load-games

add-matches: load-games
	sh ./add_match.sh
	touch add-matches

clean-db:
	docker stop hexgame-db ; docker rm hexgame-db ; \
  rm mongo-service
	rm mongo-admin-password
	rm mongo-create-hexgame-db
	rm load-games
	rm add-matches
