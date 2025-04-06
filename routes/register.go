package routes

import (
	"net/http"
	"strconv"

	"example.com/go-rest/models"
	"github.com/gin-gonic/gin"
)

func registerForEvent(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id!"})
		return
	}

	event, err := models.GetEventById(eventId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not fetch the event"})
	}

	userId := context.GetInt64("userId")
	err = event.Register(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not register for the event"})
	}
	context.JSON(http.StatusCreated, gin.H{"message": "Registered for the event"})

}

func cancelRegistration(context *gin.Context) {
	eventId, err := strconv.ParseInt(context.Param("id"), 10, 64)
	var event models.Event
	event.ID = eventId

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "Could not parse the event id!"})
		return
	}

	userId := context.GetInt64("userId")

	err = event.RemoveFromEvent(userId)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": "could not deregister from an event"})
	}
	context.JSON(http.StatusOK, gin.H{"message": "deregistered from the event!"})
}
