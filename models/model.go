package models

import (
	"time"

	"gorm.io/gorm"
)

type Book struct {
	ID          uint      `json:"id" gorm:"primaryKey"`
	Title       string    `json:"title" gorm:"column:title"`
	Author      string    `json:"author" gorm:"column:author"`
	PublishedAt time.Time `json:"published" gorm:"column:published_at;autoCreateTime"`
}

func MigrateBooks(db *gorm.DB) error {
	err := db.AutoMigrate(&Book{})
	if err != nil {
		return err
	}
	return nil
}