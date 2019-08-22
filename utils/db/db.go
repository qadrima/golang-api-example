package db

import (
	"log"

	"github.com/globalsign/mgo"
	"github.com/spf13/viper"
)

var mgoSession *mgo.Session

func InitMongo() *mgo.Session {
	if mgoSession == nil {
		var err error
		user := viper.GetString("database.mongodb.user")
		pass := viper.GetString("database.mongodb.password")
		host := viper.GetString("database.mongodb.host")
		port := viper.GetString("database.mongodb.port")
		dbUser := viper.GetString("database.mongodb.dbuser")
		url := "mongodb://" + user + ":" + pass + "@" + host + ":" + port + "/" + dbUser
		mgoSession, err = mgo.Dial(url)
		if err != nil {
			log.Fatal(err)
		}
	}
	return mgoSession
}

func Use(tableName string) (collection *mgo.Collection) {
	dbname := viper.GetString("database.mongodb.dbname")
	return mgoSession.DB(dbname).C(tableName)
}

func Close() {
	mgoSession.Close()
	return
}
