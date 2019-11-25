package routers

import (
	"github.com/gin-gonic/gin"
	"github.com/csanso/libraries/preguntas-random-api/app"
)

var router *gin.Engine

func MakeRouter(application *app.App) *gin.engine {
	mapURLs(application)
	return router
}

func mapURLs(application *app.App) {
	router.GET("/", Controller.Home)
	router.GET("/preguntas", Controller.GetPreguntas)
	router.GET("/preguntas/:id", Controller.GetPreguntaByID)
	router.POST("/preguntas", Controller.PostPregunta)
	router.DELETE("/preguntas/:id", Controller.DeletePregunta)
}


