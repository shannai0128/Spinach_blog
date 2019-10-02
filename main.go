package main

import app2 "github.com/c479096292/Spinach_blog/app"

func main() {
	app := app2.App{}
	app.InitConfig()
	app.InitRouter()
}
