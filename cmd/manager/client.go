package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"time"

	pb "github.com/mnm458/gorchestrator/api"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/protobuf/types/known/emptypb"
)

var (
	addr = flag.String("addr", "localhost:50051", "the address to connect to")
)

func main() {
	flag.Parse()
	conn, err := grpc.NewClient(*addr, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewWorkerServiceClient(conn)
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	// startResp, srtErr := c.StartTask(ctx, &pb.StartTaskRequest{
	// 	TaskEvent: &pb.TaskEvent{
	// 		Id:    []byte("266592cd-960d-4091-981c-8c25c44b1019"),
	// 		State: int32(task.Scheduled),
	// 		Task: &pb.Task{
	// 			State: int32(task.Scheduled),
	// 			Id:    []byte("266592cd-960d-4091-981c-8c25c44b1019"),
	// 			Name:  "test-chapter-5-1",
	// 			Image: "strm/helloworld-http",
	// 		},
	// 	},
	// })
	// if srtErr != nil {
	// 	log.Fatal("ERR: ", srtErr)
	// }
	// fmt.Printf("Start task response: %#v\n", startResp)
	r, err := c.GetTasks(ctx, &emptypb.Empty{})
	if err != nil {
		log.Fatal("ERR: ", err)
	}
	fmt.Printf("Response %v", r.Tasks)
}
