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
	app.GET("/preguntas", GetPreguntasEndpoint)
	app.GET("/preguntas/{id}", GetPreguntaEndpoint)
	//app.POST("/preguntas/{id}", CreatePreguntaEndpoint)
	app.DELETE("/preguntas/{id}", DeletePreguntaEndpoint)
}

// Endpoints
func Home(ctx *gin.Context) {
	ctx.String(200, "API Go funcionando v5")
	return
}
func GetPreguntasEndpoint(ctx *gin.Context) {
	ctx.JSON(200, preguntas)
	return
}
func GetPreguntaEndpoint(ctx *gin.Context) {
	param := ctx.Param("id")
	for _, pregunta := range preguntas {
		if pregunta.ID == param {
			ctx.JSON(200, pregunta)
			return
		}
	}
	ctx.JSON(200, &Pregunta{}) // Responde vacio sino encuentra pregunta
	return
}
/* func CreatePreguntaEndpoint(ctx *gin.Context) {
	param := ctx.Param("id")
	var pregunta Pregunta
	_ = json.NewDecoder(req.Body).Decode(&pregunta)
	pregunta.ID = param
	preguntas = append(preguntas, pregunta)
	json.NewEncoder(writer).Encode(preguntas)
} */
func DeletePreguntaEndpoint(ctx *gin.Context) {
	param := ctx.Param("id")
	for index, pregunta := range preguntas {
		if pregunta.ID == param {
			preguntas = append(preguntas[:index], preguntas[index+1:]...)
			ctx.String(200, "Success")
			break
		}
	}
	ctx.JSON(200, preguntas)
}
