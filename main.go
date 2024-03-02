package main

import (
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/grim13/go-rest-api/db"
	"github.com/grim13/go-rest-api/routes"
	"github.com/joho/godotenv"
)

func init() {

	err := godotenv.Load(".env")

	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":" + os.Getenv("PORT")) //localhost:8080
}
