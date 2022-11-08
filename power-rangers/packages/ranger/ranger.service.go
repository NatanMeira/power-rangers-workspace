package ranger

import (
	"github.com/gin-gonic/gin"
	db "github.com/natanmeira/power-rangers-workspace/database"
)

func index(ctx *gin.Context) {
	success := ctx.Query("success")
	deleted := ctx.Query("deleted")
	var rangers []Ranger
	db.GetDB().Order("created_at desc").Find(&rangers)
	ctx.HTML(200, "ranger.index", gin.H{
		"Rangers": rangers,
		"Success": success,
		"Deleted": deleted,
	})
}

func create(ctx *gin.Context) {
	ctx.HTML(200, "ranger.create", gin.H{
		"Disabled":   false,
		"FormAction": "/ranger/save",
		"FormLabel":  "Create",
		"FormMethod": "post",
	})
}

func insert(ctx *gin.Context) {
	var ranger Ranger
	if err := ctx.Bind(&ranger); err != nil {
		ctx.HTML(200, "ranger.create", gin.H{
			"Error": "",
		})
	}
	db.GetDB().Create(&ranger)
	ctx.Redirect(302, "/ranger?success=true")
}

func show(ctx *gin.Context) {
	id := ctx.Param("id")
	var ranger Ranger
	db.GetDB().First(&ranger, id)
	ctx.HTML(200, "ranger.show", gin.H{
		"Ranger":   ranger,
		"Disabled": true,
	})
}

func edit(ctx *gin.Context) {
	id := ctx.Param("id")
	var ranger Ranger
	db.GetDB().First(&ranger, id)
	ctx.HTML(200, "ranger.edit", gin.H{
		"Ranger":     ranger,
		"Disabled":   false,
		"FormAction": "/ranger/update/" + id,
		"FormLabel":  "Update",
		"FormMethod": "put",
	})
}

func update(ctx *gin.Context) {
	id := ctx.Param("id")
	var ranger Ranger
	db.GetDB().First(&ranger, id)
	if err := ctx.Bind(&ranger); err != nil {
		ctx.HTML(200, "ranger.edit", gin.H{
			"Error": "",
		})
	}
	db.GetDB().Model(&ranger).UpdateColumns(ranger)
	ctx.Redirect(302, "/ranger?success=true")
}

func delete(ctx *gin.Context) {
	id := ctx.Param("id")
	var ranger Ranger
	db.GetDB().Delete(&ranger, id)
	ctx.Redirect(302, "/ranger?deleted=true")
}
