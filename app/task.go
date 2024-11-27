package app

import "time"

type Status string

const (
	TODO        Status = "todo"
	IN_PROGRESS Status = "in-progress"
	DONE        Status = "done"
)

type Task struct {
	ID          int       `json:"id"`
	Description string    `json:"description"`
	Status      Status    `json:"status"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
}
