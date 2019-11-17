package main

import (
	"os"

	"github.com/gin-gonic/gin"
)

type Pregunta struct {
	ID      string `json:"id,omitempty"`
	Content string `json:"content,omitempty`
}

var preguntas []Pregunta

// Main

func main() {
	// Example data
	preguntas = append(preguntas, Pregunta{ID: "1", Content: "¿Qué comida es la que mejor te sale?"})
	preguntas = append(preguntas, Pregunta{ID: "2", Content: "¿Cuál es tu color preferido?"})

	port := os.Getenv("PORT")
	app := gin.New()
	// Routers

	app.GET("/", Home)
	app.GET("/preguntas", GetPreguntas)
	app.GET("/preguntas/:id", GetPreguntaByID)
	app.POST("/preguntas", PostPregunta)
	app.DELETE("/preguntas/:id", DeletePregunta)

	app.Run(":" + port)
}

func Home(ctx *gin.Context) {
	ctx.String(200, "API Go funcionando v6")
}

func GetPreguntas(ctx *gin.Context) {
	ctx.JSON(200, preguntas)
}

func GetPreguntaByID(ctx *gin.Context) {
	param := ctx.Param("id")
	for _, pregunta := range preguntas {
		if pregunta.ID == param {
			ctx.JSON(200, pregunta)
			return
		}
	}
	ctx.JSON(200, gin.H{
		"id":      "null",
		"Content": "null",
	})
}

func PostPregunta(ctx *gin.Context) {
	var reqBody Pregunta
	ctx.BindJSON(&reqBody)
	preguntas = append(preguntas, reqBody)
	ctx.JSON(200, preguntas)
	return
}

func DeletePregunta(ctx *gin.Context) {
	param := ctx.Param("id")
	for index, pregunta := range preguntas {
		if pregunta.ID == param {
			preguntas = append(preguntas[:index], preguntas[index+1:]...)
			ctx.JSON(200, gin.H{
				"status":  "deleted",
				"message": pregunta,
			})
			return
		}
	}
	ctx.JSON(404, gin.H{"message": "Question not found."})
}
