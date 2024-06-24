package routes

import (
	"net/http"
	"project/REST_API/models"

	"github.com/gin-gonic/gin"
)

func signup(context *gin.Context) {

	var user models.User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse request data"})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not save user", "error": err})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message": "user created successfully !"})

}

func login(context *gin.Context) {
	var user models.User

	err := context.BindJSON(&user)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = user.ValidateCredentials()
	if err != nil {
		context.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": " Logged in!"})
}
