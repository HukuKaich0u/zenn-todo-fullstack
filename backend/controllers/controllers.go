package controllers

import (
	"backend/db"
	"backend/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

func CreateTodos(c *gin.Context) {
	var input models.Input
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	todo := models.Todo{
		Title:       input.Title,
		IsCompleted: input.IsCompleted,
	}

	if err := db.DB.Create(&todo).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, todo)
}

func GetTodos(c *gin.Context) {
	var todos []models.Todo
	db.DB.Order("id asc").Find(&todos)
	c.JSON(http.StatusOK, todos)
}

func UpdateTodos(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo models.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	var input models.Todo
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	todo.IsCompleted = input.IsCompleted
	if err := db.DB.Save(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}
	c.JSON(http.StatusOK, todo)
}

func DeleteTodos(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var todo models.Todo
	if err := db.DB.First(&todo, id).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	if err := db.DB.Delete(&todo).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Todo deleted"})
}
