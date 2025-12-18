package task

import (
	"errors"
	"time"
)

type Task struct {
	ID        string
	Payload   string
	Status    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

func NewTask(id, payload string) Task {
	return Task{
		ID:        id,
		Payload:   payload,
		Status:    "pending",
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	}
}

const (
	STATUS_RUNNING   string = "running"
	STATUS_COMPLETED string = "completed"
	STATUS_FAILED    string = "failed"
)

func (t *Task) UpdateStatus(status string) error {
	if !isValidStatus(status) {
		return errors.New("invalid status")
	}
	t.Status = status
	now := time.Now()
	t.UpdatedAt = now
	return nil
}

func isValidStatus(status string) bool {
	switch status {
	case STATUS_RUNNING, STATUS_COMPLETED, STATUS_FAILED:
		return true
	default:
		return false
	}
}
