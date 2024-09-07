package worker

import (
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
}

func (w *Worker) StopTask() {

}

func (w *Worker) CollectStats() {

}
