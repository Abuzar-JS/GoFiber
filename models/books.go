package models

import (
	"gorm.io/gorm"
)

type Books struct {
	Id        int    `gorm:"primary key;autoIncrement" json:"id"`
	Author    string `gorm:"not null;unique" json:"author"`
	Title     string `gorm:"not null" json:"title"`
	Publisher string `gorm:"not null" json:"publisher"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Books{})
	return err
}
