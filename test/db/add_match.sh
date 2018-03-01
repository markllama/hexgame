HEXGAME_DATABASE=hexgame
HEXGAME_USERNAME=hexgame
HEXGAME_PASSWORD=ragnar
MONGO_ADDRESS=127.0.0.1
MONGO_LOAD="mongo --quiet --username=${HEXGAME_USERNAME} --password=${HEXGAME_PASSWORD} ${MONGO_ADDRESS}/${HEXGAME_DATABASE}"


function get_id() {
    # collection=$1
    # field=$2
    # name=$3

    echo "db.${1}.find({'$2': '$3'})" | ${MONGO_LOAD} | \
        sed -e 's/ObjectId(//' -e 's/)//' | \
        jq "._id"
}

function new_match() {
    # OWNERID=$1
    # GAMEID=$2
    #
    ${MONGO_LOAD} <<EOF
var now = new Date()
db.matches.insert(
  {
    "owner_id": ObjectId($1),
    "game_id": ObjectId($2),
    "create_time": now,
    "start_time": null,
    "players": [
      ObjectId($1)
    ],
    "map_id": null
  }
)
EOF
}

USERID=$(get_id users username mark)
GAMEID=$(get_id games name clear)
new_match $USERID $GAMEID

USERID=$(get_id users username chiara)
GAMEID=$(get_id games name blackhole)
new_match $USERID $GAMEID






