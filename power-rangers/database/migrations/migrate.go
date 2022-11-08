package migrations

import (
	db "github.com/natanmeira/power-rangers-workspace/database"
	"github.com/natanmeira/power-rangers-workspace/packages/ranger"
)

func Migrate() {
	db := db.GetDB()
	db.AutoMigrate(&ranger.Ranger{})
}
