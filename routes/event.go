package routes

import (
	"net/http"
	"strconv"

	"event.com/first/models"
	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})

	}
	context.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})

	context.JSON(http.StatusOK, events)

}

func StoreEvents(context *gin.Context) {

	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error in request", "error": err})
	}
	event.Id = 1
	event.UserId = 1
	event.StoreEvents()
	context.JSON(http.StatusCreated, gin.H{"message": "staus updated successfully", "event": event})

}

func getEventById(context *gin.Context) {
	id := context.Param("id")
	intId, err := strconv.Atoi(id)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Invalid ID"})
		return
	}

	event, err := models.GetEventById(int64(intId))
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error in request", "error": err})
	}

	context.JSON(http.StatusCreated, gin.H{"message": "get successfully", "event": event})

}
