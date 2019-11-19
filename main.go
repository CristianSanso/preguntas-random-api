package main

import (
	"os"
	"github.com/csanso/.."
	"github.com/gin-gonic/gin"
)

func main() {
	port := os.Getenv("PORT")
	api := gin.New()
	// Routers

	api.GET("/", controllers.API.Home)
	api.GET("/preguntas", controllers.API.GetPreguntas)
	api.GET("/preguntas/:id", controllers.API.GetPreguntaByID)
	api.POST("/preguntas", controllers.API.PostPregunta)
	api.DELETE("/preguntas/:id", controllers.API.DeletePregunta)

	api.Run(":" + port)
}
