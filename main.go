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

	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "API Go funcionando v5")
	})

	app.GET("/preguntas/{id}", func(ctx *gin.Context) {
		param := ctx.Param("id")
		for _, pregunta := range preguntas {
			if pregunta.ID == param {
				ctx.JSON(200, pregunta)
				return
			}
		}
		ctx.JSON(200, &Pregunta{})
	})

	app.DELETE("/preguntas/{id}", func(ctx *gin.Context) {
		param := ctx.Param("id")
		for index, pregunta := range preguntas {
			if pregunta.ID == param {
				preguntas = append(preguntas[:index], preguntas[index+1:]...)
				ctx.JSON(200, gin.H{
					"status":  "success",
					"message": pregunta,
				})
				break
			}
		}
		ctx.JSON(200, preguntas)
	})

	app.Run(":" + port)
}
