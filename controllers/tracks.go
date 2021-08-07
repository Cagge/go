// GET /tracks
// Получаем список всех треков
package controllers

import (
	"net/http"

	"github.com/alonelegion/musicstore_rest_api/models"
	"github.com/gin-gonic/gin"
)

func GetAllTracks(context *gin.Context) {
	var tracks []models.Track
	models.DB.Find(&tracks)

	context.JSON(http.StatusOK, gin.H{"tracks": tracks})
}
