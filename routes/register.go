package routes

import (
	"net/http"
	models "project/REST_API/models"
	"strconv"

	"github.com/gin-gonic/gin"
)

func regiterForEvent(context *gin.Context) {

	userId := context.GetInt64("userId")
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "could not parse the event id"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
		return
	}

	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register the event"})
		return
	}

	context.JSON(http.StatusCreated, gin.H{"message ": "event Registered successfully! :)"})

}

func getRegisteredEvents(context *gin.Context) {
	registeredEvents, err := models.GetAllRegisteredEvents()
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error message": "could not fetch Registered Events"})
		return

	}
	context.JSON(http.StatusOK, registeredEvents)
}
