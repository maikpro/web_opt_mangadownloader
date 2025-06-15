package main

import (
	"github.com/maikpro/web_opt_mangadownloader/api/routes"
	"github.com/maikpro/web_opt_mangadownloader/bootstrap"
)

func main() {
	app := bootstrap.App()
	env := app.Env

	routes.HandleHttp(env)

	//db := app.Mongo.Database(env.DBName)
	defer app.CloseDBConnection()
}
