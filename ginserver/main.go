package main

import (
	"ginserver/database"
	"ginserver/router"
)

func main() {
	database.Connect()
	app := router.SetupRouter()
	app.Run()
}
