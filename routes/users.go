package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grim13/go-rest-api/models"
)

func singup(context *gin.Context) {
	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {

		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
			"err":     err,
		})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Could not save data user.",
		})
		return
	}
	context.JSON(http.StatusCreated, gin.H{
		"message": "User created successfully.",
	})
}
