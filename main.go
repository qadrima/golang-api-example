package main

import (
	"example/config"
	"example/routes"
	"example/utils/db"

	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
)

func main() {

	// Config
	confReader := new(config.ConfigReader)
	confReader.Read()

	// DB
	dbConnect := db.InitMongo()
	defer dbConnect.Close()

	// ========== Routes init ===========
	defaultRoute := new(routes.DefaultRoute)
	router := gin.Default()
	defaultRoute.Init(router)
	router.Run(viper.GetString("server.host") + ":" + viper.GetString("server.port"))
}
