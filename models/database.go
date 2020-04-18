package models

import (
	"github.com/jinzhu/gorm"
)

var Db *gorm.DB

func AutoMigrate() {
	Db.AutoMigrate(
		&User{},
	)
}

func SetDatabase(db *gorm.DB) {
	Db = db
}
