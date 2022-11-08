package main

import (
	"fmt"

	gin "github.com/gin-gonic/gin"
	db "github.com/natanmeira/power-rangers-workspace/database"
	"github.com/natanmeira/power-rangers-workspace/database/migrations"
	route "github.com/natanmeira/power-rangers-workspace/routes"
)

func main() {

	db.Init()
	migrations.Migrate()

	router := gin.Default()

	router.Static("/assets", "./public/assets")

	router.LoadHTMLGlob("views/**/*.html")

	routesController := route.Controller{Router: router}
	routesController.InitRoutes()

	fmt.Println("Server started successfully")

	router.Run()
}
