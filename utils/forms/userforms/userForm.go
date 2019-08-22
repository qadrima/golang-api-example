package userforms

import (
	"github.com/globalsign/mgo/bson"
)

type User struct {
	Id   bson.ObjectId `json:"_id" bson:"_id"`
	Name string        `json:"name" bson:"name"`
	Age  int           `json:"age" bson:"age"`
}
