package home

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Router *gin.Engine
}

func (c *Controller) Routes() {
	c.Router.GET("/", index)
}
