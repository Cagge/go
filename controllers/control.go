package controllers

import (
	"net/http"

	"github.com/Cagge/models"
	"github.com/gin-gonic/gin"
)

type CreateAccInput struct {
	Name     string `json:"name" binding:"required"`
	Position string `json:"position" binding:"required"`
}

type UpdateAccInput struct {
	Name     string `json:"name"`
	Position string `json:"position"`
}

func GetAllAcc(context *gin.Context) {
	var accs []models.Acc
	models.DB.Find(&accs)
	context.JSON(http.StatusOK, gin.H{"accs": accs})
}

func GetAcc(context *gin.Context) {
	var acc models.Acc
	if err := models.DB.Where("id = ?", context.Param("id")).First(&acc).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Пользователя с таким ID не существует"})
		return
	}
	context.JSON(http.StatusOK, gin.H{"accs": acc})
}

func CreateAcc(context *gin.Context) {
	var input CreateAccInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Проверьте ключ и повторите ввод"})
		return
	}

	acc := models.Acc{Name: input.Name, Position: input.Position}
	models.DB.Create(&acc)
	context.JSON(http.StatusOK, gin.H{"accs": acc})
}

func UpdateAcc(context *gin.Context) {
	var acc models.Acc
	if err := models.DB.Where("id = ?", context.Param("id")).First(&acc).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Пользователя с таким ID не существует"})
		return
	}
	var input UpdateAccInput
	if err := context.ShouldBindJSON(&input); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Пользователя с таким ID не существует"})
		return
	}
	models.DB.Model(&acc).Update(input)
	context.JSON(http.StatusOK, gin.H{"acc": acc})
}

func DeleteAcc(context *gin.Context) {
	var acc models.Acc
	if err := models.DB.Where("id = ?", context.Param("id")).First(&acc).Error; err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": "Пользователя с таким ID не существует"})
		return
	}
	models.DB.Delete(&acc)
	context.JSON(http.StatusOK, gin.H{"accs": true})
}
