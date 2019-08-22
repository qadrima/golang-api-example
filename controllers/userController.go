package controllers

import (
	"example/models"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/globalsign/mgo/bson"
)

type UserController struct{}

var userModel = new(models.UserModel)

func (s *UserController) Users(c *gin.Context) {

	users, _ := userModel.GetUsers()

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   users,
	})
}

func (s *UserController) UserDetail(c *gin.Context) {

	userId := bson.ObjectIdHex(c.Param("id"))
	user, err := userModel.GetUserById(userId)

	if err != nil {
		ErrorResponse(c, http.StatusNotFound, err)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"status": http.StatusOK,
		"data":   user,
	})
}

func ErrorResponse(c *gin.Context, status int, err error) {
	c.JSON(status, gin.H{
		"status": status,
		"error":  err,
	})
	c.Abort()
}
