package types

import (
	"time"

	"gorm.io/gorm"
)

type Diary struct {
	gorm.Model
	Date    time.Time `json:"date"` // Date of what?
	Title   string    `json:"title"`
	Content string    `json:"content"`
}
