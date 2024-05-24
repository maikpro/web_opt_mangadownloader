package main

import "github.com/maikpro/web_opt_mangadownloader/routes"

// @title Web Opt Mangadownloader API
// @version 1.0
// @description This is a web api server.
// @host localhost:8080
// @BasePath /
func main() {
	routes.HandleHttp()
}
