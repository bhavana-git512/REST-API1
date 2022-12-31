package main

import (
	"manger/database"
	"manger/routes"
)

// C:\Users\jvt\go\pkg\mod\src\routes

//	"model"

func main() {

	database.Connectme()
	routes.Register()

}
