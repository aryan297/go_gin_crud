package main

import (
	"net/http"

	"event.com/first/models"
	"github.com/gin-gonic/gin"
)

func main() {

	server := gin.Default()

	server.GET("/events", getEvents)

	server.Run(":8800")

}

func getEvents(context *gin.Context) {
	events := models.GetEvents()

	/* 	context.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	}) */

	context.JSON(http.StatusOK, events)

}

func StoreEvents(context *gin.Context) {

	var event models.Event

	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Error in request"})
	}
	event.Id = 1
	context.JSON(http.StatusBadRequest, gin.H{"message": "Error in request", "event": event})

}
