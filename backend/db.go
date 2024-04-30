package main

import (
	"time"

	"gorm.io/gorm"
)

type Diary struct {
	gorm.Model
	Date    time.Time `json:"date"`
	Title   string   `json:"title"`
	Content string  `json:"content"`
}

func GetDiary() []Diary {
	diaries := []Diary{
		{
			Date:    time.Date(2024, 4, 2, 0, 0, 0, 0, time.UTC),
			Title:   "Forth Entry",
			Content: "あいうえお",
		},
		{
			Date:    time.Date(2024, 4, 3, 0, 0, 0, 0, time.UTC),
			Title:   "Fifth Entry",
			Content: "かきくけこ",
		},
	}
	return diaries
}
