package routes

import (
	"github.com/gin-gonic/gin"
	home "github.com/natanmeira/power-rangers-workspace/packages/home"
	ranger "github.com/natanmeira/power-rangers-workspace/packages/ranger"
)

type Controller struct {
	Router *gin.Engine
}

func (c *Controller) InitRoutes() {

	c.Router.GET("/healthz", func(c *gin.Context) {
		c.JSON(200, "ok")
	})

	homeController := home.Controller{Router: c.Router}
	homeController.Routes()
	rangerController := ranger.Controller{Router: c.Router}
	rangerController.Routes()
}
