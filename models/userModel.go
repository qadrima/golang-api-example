package models

import (
	"example/utils/db"
	"example/utils/forms/userforms"

	"github.com/globalsign/mgo/bson"
)

type UserModel struct{}

func (m *UserModel) GetUsers() (users []userforms.User, err error) {
	collection := db.Use("users")
	err = collection.Find(bson.M{}).All(&users)
	return users, err
}

func (m *UserModel) GetUserById(id bson.ObjectId) (user userforms.User, err error) {
	collection := db.Use("users")
	err = collection.Find(bson.M{"_id": id}).One(&user)
	return user, err
}
