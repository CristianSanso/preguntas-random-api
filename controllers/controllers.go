
package controllers

type API struct {}

func (self *API) Home(ctx *gin.Context) {
	ctx.String(200, "API Go funcionando v6")
}

func (self *API) GetPreguntas(ctx *gin.Context) {
	ctx.JSON(200, preguntas)
}

func (self *API) GetPreguntaByID(ctx *gin.Context) {
	param := ctx.Param("id")
	for _, pregunta := range preguntas {
		if pregunta.ID == param {
			ctx.JSON(200, pregunta)
			return
		}
	}
	ctx.JSON(200, gin.H{
		models.Pregunta{ID: "null", Content: "null"}
	})
}

func (self *API) PostPregunta(ctx *gin.Context) {
	var reqBody Pregunta
	ctx.BindJSON(&reqBody)
	preguntas = append(preguntas, reqBody)
	ctx.JSON(200, preguntas)
	return
}

func (self *API) DeletePregunta(ctx *gin.Context) {
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
