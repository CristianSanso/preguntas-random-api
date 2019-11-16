package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/mux"
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
	app.GET("/", home)
	app.GET("/preguntas", GetPreguntasEndpoint)
	app.GET("/preguntas/{id}", GetPreguntaEndpoint)

	app.POST("/preguntas/{id}", CreatePreguntaEndpoint)

	app.DELETE("/preguntas/{id}", DeletePreguntaEndpoint)

	log.Fatal(http.ListenAndServe(":3000", router))
}

// Endpoints

func home(ctx *gin.Context) {
	ctx.String(200, "API Go funcionando v4")
}

func GetPreguntasEndpoint(ctx *gin.Context) {
	ctx.JSON(200, preguntas)
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
}

func CreatePreguntaEndpoint(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	var pregunta Pregunta
	_ = json.NewDecoder(req.Body).Decode(&pregunta)
	pregunta.ID = params["id"]
	preguntas = append(preguntas, pregunta)
	json.NewEncoder(writer).Encode(preguntas)
}

func DeletePreguntaEndpoint(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for index, pregunta := range preguntas {
		if pregunta.ID == params["id"] {
			preguntas = append(preguntas[:index], preguntas[index+1:]...)
			break
		}
	}
	json.NewEncoder(writer).Encode(preguntas)
}
