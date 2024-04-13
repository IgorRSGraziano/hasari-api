package models

import (
	"time"

	"gorm.io/gorm"
)

type Music struct {
	gorm.Model
	ID      int       `json:"id"`
	Title   string    `json:"title"`
	Audio   []byte    `json:"audio"`
	Release time.Time `json:"release"`
	AlbumID int       `json:"albumId"`
	Album   Album     `json:"album"`
}
