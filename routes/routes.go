package routes

import (
	"project/REST_API/middleware"

	"github.com/gin-gonic/gin"
)

func RegisterRoutes(server *gin.Engine) {

	server.GET("/events", getEvents)
	server.GET("/event/:id", getEventById)

	// Routes with Auth Middleware
	authenticated := server.Group("/")
	authenticated.Use(middleware.Authenticate)
	authenticated.POST("/events", createEvent)
	authenticated.PUT("/update/:id", updateEvent)
	authenticated.DELETE("/delete/:id", deleteEvent)

	server.POST("/signup", signup)
	server.POST("/login", login)
}
