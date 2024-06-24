package routes

import (
	"net/http"
	models "project/REST_API/events"
	"strconv"

	"github.com/gin-gonic/gin"
)

func getEvents(context *gin.Context) {
	events, err := models.GetAllEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error message": "could not fetch events"})
		return

	}

	context.JSON(http.StatusOK, events)

}

func getEventById(ctx *gin.Context) {
	eventId, err := strconv.ParseInt(ctx.Param("id"), 10, 64)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the event id"})
		return
	}
	event, err := models.GetEventById(eventId)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event by ID"})
		return
	}
	ctx.JSON(http.StatusOK, event)
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

	err = event.Save()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error message": "saving event failed"})
		return

	}

	context.JSON(http.StatusCreated, gin.H{"message ": "event created :)", "event": event})

}

func deleteAllEvents(ctx *gin.Context) {
	err := models.DeleteAllEvents()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error message": "could not delete all rows"})
		return
	} else {
		ctx.JSON(http.StatusInternalServerError, gin.H{"error message": "All rows are deleted"})
	}
}
