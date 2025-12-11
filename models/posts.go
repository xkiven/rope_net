package models

import "time"

type Post struct {
	ID          uint
	Title       string
	Content     string
	UserID      uint
	PageView    int
	PublishTime time.Time
}
