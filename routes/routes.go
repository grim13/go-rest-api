package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/grim13/go-rest-api/middlewares"
)

func RegisterRoutes(server *gin.Engine) {
	server.GET("/events", getEvents) //GET, POST, PUT, PATCH, DELETE
	server.GET("/events/:id", getEvent)

	autenticate := server.Group("/")
	autenticate.Use(middlewares.Authenticate)
	autenticate.POST("/events", createEvent)
	autenticate.PUT("/events/:id", updateEvent)
	autenticate.DELETE("/events/:id", deleteEvent)

	server.POST("/signup", singup)
	server.POST("/login", login)
}
