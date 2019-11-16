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
	port := os.Getenv("PORT")
	app := gin.New()
	app.GET("/", func(ctx *gin.Context) {
		ctx.String(200, "API Go funcionando v2")
	})
	app.Run(":" + port)

	router := mux.NewRouter()

	// Example data
	preguntas = append(preguntas, Pregunta{ID: "1", Content: "¿Qué comida es la que mejor te sale?"})
	preguntas = append(preguntas, Pregunta{ID: "2", Content: "¿Cuál es tu color preferido?"})

	// Endpoints
	app.GET("/preguntas", func(ctx *gin.Context) {
		ctx.JSON(200, preguntas)
	})
	router.HandleFunc("/preguntas/{id}", GetPreguntaEndpoint).Methods("GET")
	router.HandleFunc("/preguntas/{id}", CreatePreguntaEndpoint).Methods("POST")
	router.HandleFunc("/preguntas/{id}", DeletePreguntaEndpoint).Methods("DELETE")

	log.Fatal(http.ListenAndServe(":3000", router))
}

// Endpoints
func GetPreguntaEndpoint(writer http.ResponseWriter, req *http.Request) {
	params := mux.Vars(req)
	for _, pregunta := range preguntas {
		if pregunta.ID == params["id"] {
			json.NewEncoder(writer).Encode(pregunta)
			return
		}
	}
	json.NewEncoder(writer).Encode(&Pregunta{}) // Responde vacio sino encuentra pregunta
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
