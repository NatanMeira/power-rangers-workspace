package home

import (
	"fmt"

	"github.com/gin-gonic/gin"
	db "github.com/natanmeira/power-rangers-workspace/database"
	r "github.com/natanmeira/power-rangers-workspace/packages/ranger"
)

func index(ctx *gin.Context) {
	var rangers []r.Ranger
	db.GetDB().Limit(6).Order("created_at desc").Find(&rangers)
	fmt.Println(rangers)
	ctx.HTML(200, "home", gin.H{
		"Rangers": rangers,
		"Title":   "Home | Power Rangers",
	})
}
