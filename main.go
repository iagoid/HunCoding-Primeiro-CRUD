package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/iagoid/HunCoding-Primeiro-CRUD/src/controller/routes"
	"github.com/joho/godotenv"
)

func main() {

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)
	err = router.Run(":8080")
	if err != nil {
		log.Fatal(err)
	}
}
