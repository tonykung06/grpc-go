package main

import (
	"fmt"
	"io"

	"github.com/grpc-go/pb"

	"google.golang.org/grpc/metadata"

	"golang.org/x/net/context"
)

type employeeServer struct {
}

func (es *employeeServer) GetByBadgeNumber(ctx context.Context,
	req *pb.GetByBadgeNumberRequest) (*pb.EmployeeResponse, error) {
	if md, ok := metadata.FromContext(ctx); ok {
		fmt.Printf("Metadata received: %v\n", md)
	}
	employee, err := GetEmployeeByBadgeNumber(int(req.BadgeNumber))
	if err != nil {
		return nil, err
	}
	response := &pb.EmployeeResponse{Employee: convertEmployeeToMessage(employee)}
	return response, nil
}

func (es *employeeServer) GetAll(req *pb.GetAllRequest,
	stream pb.EmployeeService_GetAllServer) error {
	employees, err := GetAllEmployees()
	if err != nil {
		return err
	}
	for _, emp := range employees {
		if err = stream.Send(&pb.EmployeeResponse{Employee: convertEmployeeToMessage(emp)}); err != nil {
			return err
		}
	}
	return nil
}

func (es *employeeServer) Save(ctx context.Context,
	req *pb.EmployeeRequest) (*pb.EmployeeResponse, error) {
	emp := convertMessageToEmployee(req.Employee)
	emp, err := SaveEmployee(emp)
	if err != nil {
		return nil, err
	}
	response := &pb.EmployeeResponse{Employee: convertEmployeeToMessage(emp)}
	return response, nil
}

func (es *employeeServer) SaveAll(stream pb.EmployeeService_SaveAllServer) error {
	counter := 0
	for {
		empMsg, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("%v entities are saved.\n", counter)
			return nil
		}
		if err != nil {
			return err
		}
		emp := convertMessageToEmployee(empMsg.Employee)
		emp, err = SaveEmployee(emp)
		if err != nil {
			return err
		}
		response := convertEmployeeToMessage(emp)
		err = stream.Send(&pb.EmployeeResponse{Employee: response})

		if err != nil {
			return err
		}
		counter++
	}
}

func (es *employeeServer) AddPhoto(stream pb.EmployeeService_AddPhotoServer) error {
	md, ok := metadata.FromContext(stream.Context())
	if ok {
		fmt.Printf("Receiving photo for badge number %v\n", md["badgenumber"][0])
	}
	imgData := []byte{}
	for {
		data, err := stream.Recv()
		if err == io.EOF {
			fmt.Printf("File received with length: %v\n", len(imgData))
			return stream.SendAndClose(&pb.AddPhotoResponse{IsOk: true})
		}
		if err != nil {
			return err
		}
		fmt.Printf("Received %v bytes\n", len(data.Data))
		imgData = append(imgData, data.Data...)
	}
}
