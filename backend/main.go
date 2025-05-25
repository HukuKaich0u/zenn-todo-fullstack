package main

import (
	"backend/controllers"
	"backend/db"
	"backend/middleware"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	db.ConnectDB()

	allowOrigins := []string{"http://localhost:3000"}
	router.Use(middleware.CorsMiddleware(allowOrigins))

	router.GET("/todos", controllers.GetTodos)

	router.Run()
}
