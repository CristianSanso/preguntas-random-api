package app

import (
	"github.com/csanso/libraries/preguntas-random-api/controllers"
)

type App struct {
	Controller controllers.Controller
}

func BuildApp() *App {
	return &App
}