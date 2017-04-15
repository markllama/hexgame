db = db.getSiblingDB('hexgame')
db.createUser({user: "hexgame", pwd: "ragnar", roles: [ "readWrite", "dbAdmin" ]})



