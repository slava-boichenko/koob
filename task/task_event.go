package task

import (
	"time"

	"koob/job"

	"github.com/google/uuid"
)

type TaskEvent struct {
	ID        uuid.UUID
	Task      *Task
	Timestamp time.Time
	State     job.State
}
