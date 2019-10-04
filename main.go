package main

import (
	app2 "github.com/c479096292/Spinach_blog/app"
	"github.com/c479096292/Spinach_blog/model"
)

func main() {

	app := app2.App{}
	app.InitConfig()
	model.InitRelations()

	app.InitRouter()
}
