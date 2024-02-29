package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grim13/go-rest-api/models"
	"github.com/grim13/go-rest-api/requests"
	"github.com/grim13/go-rest-api/utils"
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

func login(context *gin.Context) {
	var req requests.Login
	err := context.ShouldBindJSON(&req)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{
			"message": "Could not parse request data.",
			"err":     err,
		})
		return
	}

	user := models.User{
		Email:    req.Email,
		Password: req.Password,
	}

	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{
			"message": "Authication Failed!!!!",
		})
		return
	}

	token, err := utils.GenerateToken(user.Email, user.ID)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{
			"message": "Authication Failed!",
		})
		return
	}

	context.JSON(http.StatusOK, gin.H{"messege": "Login Succesfull", "token": token})

}
