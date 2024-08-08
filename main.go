package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/controller"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/controller/routes"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/model/service"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init dependencies
	userService := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(userService)

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup, userController)
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
