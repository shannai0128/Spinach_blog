package app

import (
	"github.com/c479096292/Spinach_blog/config"
	"github.com/c479096292/Spinach_blog/db"
	"github.com/c479096292/Spinach_blog/router"
	"github.com/c479096292/Spinach_blog/utils"
)

type App struct {
	//DB *sqlx.DB
	//common.Response
}

func (app *App) InitRouter() {
	router.LoadRouterHandler()
}

func (app *App) InitConfig() {

	file := config.ParseConfigInfo(2)

	go utils.CheckLogSize(file) // 日志监控

	db.InitDB()
	db.InitRedis()
}

