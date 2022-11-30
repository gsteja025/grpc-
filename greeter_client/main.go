package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "example.com/helloworld/helloworld"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	addr = "localhost:50051"
)

func main() {
	conn, err := grpc.Dial(addr, grpc.WithTransportCredentials(insecure.NewCredentials()), grpc.WithBlock())
	if err != nil {
		log.Fatal("there is this error.....", err.Error())
	}
	defer conn.Close()
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	c := pb.NewEmployeecrudClient(conn)
	emp_new, err := c.CreateEmployee(ctx, &pb.Resemp{Name: "akshith", DepartmentId: 6})

	if err != nil {
		log.Fatal(err.Error())
	}

	log.Printf("id  %v, name  %v", emp_new.GetId(), emp_new.GetName())
	emps, err := c.GetEmployees(ctx, &pb.VoidEmpRequest{})
	if err != nil {
		log.Printf("connection error")
	}
	for _, i := range emps.Employees {
		fmt.Println(i.GetName(), i.GetId())
	}

}
