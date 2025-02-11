package models

import "time"

type ThreadComments struct {
	ID         uint
	UserID     uint
	CommentID  uint
	Content    string
	CreateTime time.Time
}
