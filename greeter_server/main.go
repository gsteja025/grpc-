package main

import (
	"context"
	"log"
	"net"

	pb "example.com/helloworld/helloworld"
	model "example.com/helloworld/mydb"
	"github.com/jinzhu/gorm"
	_ "github.com/lib/pq"
	"google.golang.org/grpc"
)

const (
	port = ":50051"
)

func (ptr *myDBserver) CreateEmployee(ctx context.Context, in *pb.Resemp) (*pb.Emp, error) {
	emp := model.Employee{
		Name:         in.GetName(),
		DepartmentID: uint(in.GetDepartmentId()),
	}
	ptr.db.Save(&emp)
	log.Printf("%v %v %v", in.GetName(), in.GetDepartmentId())
	return &pb.Emp{Name: in.GetName(), DepartmentId: int64(in.GetDepartmentId()), Id: int64(emp.ID)}, nil
}

func (ptr *myDBserver) GetEmployees(ctx context.Context, in *pb.VoidEmpRequest) (*pb.Employees, error) {
	empdata := []model.Employee{}
	ptr.db.Find(&empdata)

	Res := []*pb.Emp{}
	for _, val := range empdata {
		Res = append(Res, &pb.Emp{Name: val.Name, DepartmentId: int64(val.DepartmentID), Id: int64(val.ID)})
	}

	return &pb.Employees{Employees: Res}, nil
}

func (ptr *myDBserver) DeleteEmployee(ctx context.Context, in *pb.Emp) (*pb.VoidEmpResponse, error) {

	emp := &model.Employee{
		Name: in.GetName(),
	}
	ref := &model.Employee{}
	ptr.db.Where(emp).Delete(ref)
	return &pb.VoidEmpResponse{}, nil
}

func main() {

	conn, err := gorm.Open("postgres", "user=postgres password=root dbname=postgres sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	defer conn.Close()
	lis, err := net.Listen("tcp", port)
	log.Printf("Listening to server on %v", lis.Addr())
	ser := grpc.NewServer()

	pb.RegisterEmployeecrudServer(ser, &myDBserver{
		db: conn,
	})

	if err := ser.Serve(lis); err != nil {
		log.Fatal("cant handle this", err.Error())
	}

}

type myDBserver struct {
	pb.UnimplementedEmployeecrudServer
	db *gorm.DB
}
