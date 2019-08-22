[Go](https://golang.org/)

Instalation: 
1. Make sure Go already [installed](https://golang.org/doc/install) in your machine and setup your [Go Workspace](https://golang.org/doc/code.html#Workspaces)
2. Install [Govendor](https://github.com/kardianos/govendor) `go get github.com/kardianos/govendor`
3. Go to your $GOPATH/src/. Run `git clone https://github.com/qadrima/golang-api-example`
4. Go to inside $GOPATH/src/golang-api-example
5. Run `govendor sync` to pull all package needed and wait until finish
6. Adjust your config.json file
7. Install [Fresh](https://github.com/gravityblast/fresh)
8. You're ready to go. Run `go run main.go`
9. Now your application is running in your specific port in config

*Dont forget to change mode to `production` in config/config.json

Depedencies :
- [Gin](https://github.com/gin-gonic/gin)
- [Viper](https://github.com/spf13/viper)
- [GORM](https://github.com/jinzhu/gorm)
- [MongoDB](https://github.com/globalsign/mgo)
