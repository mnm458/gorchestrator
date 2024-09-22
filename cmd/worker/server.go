package main

import (
	"context"
	"errors"
	"flag"

	"github.com/google/martian/log"
	"github.com/google/uuid"
	pb "github.com/mnm458/gorchestrator/api"
	"github.com/mnm458/gorchestrator/task"
	"github.com/mnm458/gorchestrator/worker"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	port = flag.Int("port", 50051, "server port")
)

type WorkerServer struct {
	pb.UnimplementedWorkerServiceServer
	worker *worker.Worker
}

func (ws *WorkerServer) StartTask(ctx context.Context, req *pb.StartTaskRequest) (*pb.StartTaskResponse, error) {
	var taskEvent task.TaskEvent
	taskEvent.ID = uuid.UUID(req.TaskEvent.Id)
	taskEvent.State = task.State(req.TaskEvent.State)
	taskEvent.Timestamp = req.TaskEvent.Timestamp.AsTime()
	t := task.Task{
		ID:            uuid.UUID(req.TaskEvent.Task.Id),
		ContainerID:   req.TaskEvent.Task.ContainerId,
		Name:          req.TaskEvent.Task.Name,
		State:         task.State(req.TaskEvent.State),
		Image:         req.TaskEvent.Task.Image,
		Memory:        int(req.TaskEvent.Task.Memory),
		Disk:          int(req.TaskEvent.Task.Disk),
		RestartPolicy: req.TaskEvent.Task.RestartPolicy,
	}
	ws.worker.AddTask(t)
	log.Info("Task added")
	return &pb.StartTaskResponse{
		Task: &pb.Task{
			Id:            t.ID[:],
			ContainerId:   t.ContainerID,
			Name:          t.Name,
			State:         int32(t.State),
			Image:         t.Image,
			Memory:        int32(t.Memory),
			Disk:          int32(t.Disk),
			RestartPolicy: t.RestartPolicy,
		},
	}, nil
}

func (ws *WorkerServer) GetTasks(ctx context.Context, req *emptypb.Empty) (*pb.GetTasksResponse, error) {
	tasks := ws.worker.GetTasks()
	protoTasks := packTasks(tasks)
	return &pb.GetTasksResponse{
		Tasks: protoTasks,
	}, nil
}

/*
Okay, so now we have a taskID and we have converted it to the correct type. The next thing we want to do is to check if
the worker actually knows about this task. If it doesn’t, then we should return a response with a 404 status code.
If it does, then we change the state to task.Completed and add it to the worker’s queue.
This is what the remaining of the method is doing.
*/
func (ws *WorkerServer) StopTask(ctx context.Context, req *pb.StopTaskRequest) (*emptypb.Empty, error) {
	taskID := uuid.UUID(req.Id)
	if _, exists := ws.worker.Db[taskID]; !exists {
		return &emptypb.Empty{}, errors.New("task ID does not exist")
	}
	t := ws.worker.Db[taskID]
	tCopy := *t
	tCopy.State = task.Completed
	ws.worker.AddTask(tCopy)
	return &emptypb.Empty{}, nil
}

//--- H E L P E R S

func packTasks(tasks []*task.Task) []*pb.Task {
	var res []*pb.Task
	for _, t := range tasks {
		protoTask := &pb.Task{
			Id:            t.ID[:],
			ContainerId:   t.ContainerID,
			Name:          t.Name,
			State:         int32(t.State),
			Image:         t.Image,
			Memory:        int32(t.Memory),
			Disk:          int32(t.Disk),
			RestartPolicy: t.RestartPolicy,
		}
		res = append(res, protoTask)
	}
	return res
}
