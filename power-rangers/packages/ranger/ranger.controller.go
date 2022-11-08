package ranger

import (
	"github.com/gin-gonic/gin"
)

type Controller struct {
	Router *gin.Engine
}

func (c *Controller) Routes() {
	routerGroup := c.Router.Group("/ranger")
	routerGroup.GET("/", index)
	routerGroup.GET("/new", create)
	routerGroup.POST("/save", insert)
	routerGroup.GET("/show/:id", show)
	routerGroup.GET("/edit/:id", edit)
	routerGroup.POST("/update/:id", update)
	routerGroup.GET("/delete/:id", delete)
}
