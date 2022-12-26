package app

import (
	"database/sql"
	"main/controllers"
	"main/repositories"
	"main/services"

	"github.com/gin-gonic/gin"
	_ "github.com/lib/pq"
)

func App(databaseUrl string, port string) {
	db, _ := sql.Open("postgres", databaseUrl)

	usersRepo := repositories.NewUsersRepository(db)
	usersService := services.NewUsersService(usersRepo)
	usersController := controllers.NewUsersController(usersService)
	
	router := gin.Default()

	router.GET("/health", controllers.Health)
	router.GET("/users/:user_id", usersController.Get)
	router.POST("/users", usersController.Create)

	router.Run(":" + port)
}