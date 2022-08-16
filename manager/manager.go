package manager

import (
	"fmt"
	"koob/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Manager struct {
	Pending       queue.Queue
	TaskDb        map[uuid.UUID]task.Task
	EventDb       map[uuid.UUID]task.TaskEvent
	Workers       []string
	WorkerTaskMap map[string][]uuid.UUID
	TaskWorkerMap map[uuid.UUID]string
}

func (m *Manager) SelectWorker() {
	fmt.Println("manager select worker")
}

func (m *Manager) UpdateTasks() {
	fmt.Println("manager update tasks")
}

func (m *Manager) SendWork() {
	fmt.Println("manager send work")
}
