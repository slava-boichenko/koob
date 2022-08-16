package task

import (
	"koob/job"
	"time"

	"github.com/docker/go-connections/nat"
	"github.com/google/uuid"
)

type Task struct {
	ID            uuid.UUID
	Name          string
	State         job.State
	Image         string
	Memory        int
	Disk          int
	ExposedPorts  nat.PortSet
	PortBindings  map[string]string
	RestartPolicy string
	StartTime     time.Time
	FinishTime    time.Time
}
