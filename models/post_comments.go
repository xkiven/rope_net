package models

import "time"

type PostComment struct {
	ID         uint
	UserID     uint
	PostID     uint
	Content    string
	CreateTime time.Time
}
