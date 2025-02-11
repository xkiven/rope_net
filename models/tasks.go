package models

import "time"

type Task struct {
	ID        uint
	UserID    uint
	Name      string
	Deadline  time.Time
	Completed bool
}
