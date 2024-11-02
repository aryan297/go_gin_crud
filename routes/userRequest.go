package routes

import (
	"net/http"

	"event.com/first/models"
	"github.com/gin-gonic/gin"
)

func SignUp(context *gin.Context) {

	var user models.User

	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "bad request", "error": err})
		return
	}

	err = user.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "Internal Server error", "error": err})
		return
	}
	context.JSON(http.StatusOK, gin.H{"message": "user created successfully"})

}
