package main

import (
	"github.com/gin-gonic/gin"
	"github.com/grim13/go-rest-api/db"
	"github.com/grim13/go-rest-api/routes"
)

func main() {
	db.InitDB()
	server := gin.Default()

	routes.RegisterRoutes(server)

	server.Run(":8080") //localhost:8080
}
