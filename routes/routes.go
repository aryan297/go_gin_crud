package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.POST("store/events", StoreEvents)
	server.GET("/events/:id", getEventById)
	server.PUT("events/:id", updateEvents)

}
