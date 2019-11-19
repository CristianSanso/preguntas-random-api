package routers

import (
	"github.com/gin-gonic/gin"
)

api := gin.New()

api.GET("/", Home)
api.GET("/preguntas", GetPreguntas)
api.GET("/preguntas/:id", GetPreguntaByID)
api.POST("/preguntas", PostPregunta)
api.DELETE("/preguntas/:id", DeletePregunta)

api.Run(":" + port)
