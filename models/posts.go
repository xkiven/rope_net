package models

import "time"

type Post struct {
	Title       string
	Content     string
	UserID      uint
	PageView    int
	PublishTime time.Time
}
