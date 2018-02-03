package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)


type User struct {
	Username string `json:"username"`
	FirstName string `json:"first_name"`
	Surname string `json:"surname"`
	EmailAddress string `json:"email_address"`
}

func (u *User) Get(c *mgo.Collection, selector bson.M) (error) {
	return c.Find(selector).One(&u)
}
