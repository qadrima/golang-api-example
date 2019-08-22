package routes

import (
	"example/controllers"

	"github.com/gin-gonic/gin"
)

type DefaultRoute struct{}

// controllers :
var exampleController = new(controllers.ExampleController)
var userController = new(controllers.UserController)

func (d *DefaultRoute) Init(router *gin.Engine) {

	// ========== Routers ===========
	v1 := router.Group("/v1")
	{
		v1.GET("/ping", exampleController.Index)

		// user
		user := v1.Group("/user")
		{
			user.GET("/", userController.Users)
			user.GET("/:id", userController.UserDetail)
		}

	}

	router.NoRoute(func(c *gin.Context) {
		var d = struct {
			Status  int    `json:"status"`
			Message string `json:"message"`
		}{
			Status:  404,
			Message: "Url not found.",
		}
		c.JSON(404, d)
	})
}
