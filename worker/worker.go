package worker

import (
	"fmt"
	"koob/task"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
)

type Worker struct {
	Name      string
	Db        map[uuid.UUID]task.Task
	Queue     queue.Queue
	TaskCount int
}

func (w *Worker) CollectStats() {
	fmt.Println("worker collect stats")
}

func (w *Worker) RunTask() {
	fmt.Println("worker run task")
}

func (w *Worker) StartTask() {
	fmt.Println("worker start task")
}

func (w *Worker) StopTask() {
	fmt.Println("worker stop task")
}
