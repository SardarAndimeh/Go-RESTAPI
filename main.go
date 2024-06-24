package main

import (
<<<<<<< HEAD
	"net/http"
	models "project/REST_API/events"
=======
	"project/REST_API/routes"
>>>>>>> 8e986f10914499cf35aed1d00ed73fe98c9c2b56

	"project/REST_API/db"

	"github.com/gin-gonic/gin"
)

func main() {

	db.InitDB()
	server := gin.Default()

<<<<<<< HEAD
	server.GET("/events", getEvents)
	server.POST("/events", postEvents)
=======
	routes.RegisterRoutes(server)
>>>>>>> 8e986f10914499cf35aed1d00ed73fe98c9c2b56

	server.Run(":8080")

}
<<<<<<< HEAD

func getEvents(context *gin.Context) {
	events := models.GetAllEvents()
	context.JSON(http.StatusOK, events)
}

func postEvents(context *gin.Context) {

	var event models.Event
	err := context.ShouldBindJSON(&event)

	if err != nil {
		context.JSON(http.StatusBadRequest, err)
		return
	}

	event.ID = 1
	event.UserId = 1

	event.Save()

	context.JSON(http.StatusCreated, gin.H{"message ": "event created :)", "event": event})

}
=======
>>>>>>> 8e986f10914499cf35aed1d00ed73fe98c9c2b56
