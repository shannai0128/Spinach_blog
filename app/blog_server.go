package app

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
)

type App struct {
	DB *sqlx.DB
	route *gin.Engine
}

func (app *App) RunApp() {
	route := gin.Default()
	route.Run("localhost:9000")
}



