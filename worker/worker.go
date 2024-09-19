package worker

import (
	"fmt"

	"github.com/golang-collections/collections/queue"
	"github.com/google/uuid"
	"github.com/mnm458/gorchestrator/task"
)

type Worker struct {
	Name      string
	Queue     queue.Queue
	Db        map[uuid.UUID]task.Task
	TaskCount int
}

func (w *Worker) RunTask() error {
	return nil
}

func (w *Worker) StartTask() {
	fmt.Println("StartTask called")

}

func (w *Worker) StopTask() {
	fmt.Println("StopTask called")

}

func (w *Worker) CollectStats() {
	fmt.Println("CollectStats called")

}
