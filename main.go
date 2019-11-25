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
	app.PUT("/preguntas/:id", PutPregunta)
	app.DELETE("/preguntas/:id", DeletePregunta)

	app.Run(":" + port)
}

func Home(ctx *gin.Context) {
	ctx.String(200, `
	Preguntas API V3
	Base URL: https://preguntas-random.herokuapp.com/

	GET    /preguntas 	   list all preguntas
	GET    /preguntas/:id  return pregunta from id
	POST   /preguntas	   create new pregunta
	PUT    /preguntas/:id  modify existing pregunta
	DELETE /preguntas/:id  delete existing pregunta

	The body in POST, PUT and DELETE cases must be:
	{
		"Id": "{id}",
		"Content": "{content}"
	}

	Good Luck 😄.
	`)
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

func setPregunta(id string, pregunta Pregunta) bool {
	var preguntasCopy []Pregunta
	var err bool = true
	for _, element := range preguntas {
		if element.ID == id {
			err = false
			element.ID = pregunta.ID
			println("element content", element.Content)
			element.Content = pregunta.Content
			println("pregunta content", pregunta.Content)
		}
		preguntasCopy = append(preguntasCopy, element)
	}
	if !err {
		preguntas = preguntasCopy
	}
	return err
}

func findID(id string) bool {
	var findID bool = false
	for _, element := range preguntas {
		if element.ID == id {
			findID = true
		}
	}
	return findID
}

func checkPregunta(pregunta Pregunta) bool {
	id := pregunta.ID
	content := pregunta.Content
	existID := findID(id)
	if content != "" && content != " " && existID == false {
		return true
	}
	return false
}

func PostPregunta(ctx *gin.Context) {
	var reqBody Pregunta
	ctx.BindJSON(&reqBody)
	if checkPregunta(reqBody) {
		preguntas = append(preguntas, reqBody)
		ctx.JSON(200, preguntas)
		return
	}
	ctx.JSON(400, "The request or the syntax is invalid, sonso.")
	return
}

func PutPregunta(ctx *gin.Context) {
	param := ctx.Param("id")
	var reqBody Pregunta
	ctx.BindJSON(&reqBody)
	err := setPregunta(param, reqBody)
	if !err {
		ctx.JSON(200, preguntas)
		return
	}
	ctx.JSON(400, "The request or the syntax is invalid, sonso.")
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