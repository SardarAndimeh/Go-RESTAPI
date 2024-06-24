package routes

import "github.com/gin-gonic/gin"

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/event/:id", getEventById)
	server.GET("/deleteAll", deleteAllEvents)
	server.POST("/events", postEvents)
	server.PUT("/update/:id", updateEvent)
}
