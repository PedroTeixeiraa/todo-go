package models

import "time"

type Task struct {
	ID          int       `db:"id" json:"id"`
	Title       string    `db:"title" json:"title"`
	Description string    `db:"description" json:"description"`
	Status      string    `db:"status" json:"status"` // e.g., "pending", "completed"
	DueDate     time.Time `db:"due_date" json:"due_date"`
}
