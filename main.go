package main

import (
	"os"
	"github.com/gin-gonic/gin"
	"github.com/csanso/libraries/preguntas-random-api/app"
	"github.com/csanso/libraries/preguntas-random-api/routers"
)

type App struct {
	APIController controllers.Controller
}

func main() {
	port := os.Getenv("PORT")

	application := app.BuildApp()
	router := routers.MakeRouter(application)

	router.Run(":" + port)
}
