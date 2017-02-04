package main

import (
	"log"
	"path"
	"path/filepath"
	"time"

	"fmt"

	"flag"
	"io"
	"os"

	"github.com/grpc-go/pb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/metadata"
)

const port = ":9000"

func main() {
	option := flag.Int("o", 1, "Command to run")
	flag.Parse()
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	creds, err := credentials.NewClientTLSFromFile(path.Join(dir, "cert.pem"), "")
	if err != nil {
		log.Fatal(err)
	}
	// creds := credentials.NewTLS(&tls.Config{InsecureSkipVerify: true})
	opts := []grpc.DialOption{grpc.WithTransportCredentials(creds)}
	conn, err := grpc.Dial("localhost"+port, opts...)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	client := pb.NewEmployeeServiceClient(conn)
	switch *option {
	case 1:
		SendMetadata(client)
	case 2:
		GetByBadgeNumber(client)
	case 3:
		GetAll(client)
	case 4:
		AddPhoto(client)
	case 5:
		SaveAll(client)
	}
}

func SendMetadata(client pb.EmployeeServiceClient) {
	md := metadata.MD{}
	md["user"] = []string{"mvansickle"}
	md["password"] = []string{"password1"}
	ctx := context.Background()
	ctx = metadata.NewContext(ctx, md)
	client.GetByBadgeNumber(ctx,
		&pb.GetByBadgeNumberRequest{})
}

func GetByBadgeNumber(client pb.EmployeeServiceClient) {
	res, _ := client.GetByBadgeNumber(context.Background(),
		&pb.GetByBadgeNumberRequest{BadgeNumber: 64927})

	fmt.Println(res)
}

func GetAll(client pb.EmployeeServiceClient) {
	stream, err := client.GetAll(context.Background(), &pb.GetAllRequest{})
	if err != nil {
		log.Fatal(err)
	}
	for {
		res, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(res.Employee)
	}
}

func AddPhoto(client pb.EmployeeServiceClient) {
	dir, err := filepath.Abs(filepath.Dir(os.Args[0]))
	if err != nil {
		log.Fatal(err)
	}
	f, err := os.Open(path.Join(dir, "Penguins.jpg"))
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	md := metadata.New(map[string]string{"badgenumber": "2080"})
	ctx := context.Background()
	ctx = metadata.NewContext(ctx, md)
	stream, err := client.AddPhoto(ctx)
	if err != nil {
		log.Fatal(err)
	}
	for {
		chunk := make([]byte, 64*1024)
		n, err := f.Read(chunk)
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		if n < len(chunk) {
			chunk = chunk[:n]
		}
		stream.Send(&pb.AddPhotoRequest{Data: chunk})
	}
	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(res.IsOk)
}

func SaveAll(client pb.EmployeeServiceClient) {
	employees := []pb.Employee{
		pb.Employee{
			BadgeNumber:         123,
			FirstName:           "John",
			LastName:            "Smith",
			VacationAccrualRate: 1.2,
			VacationAccrued:     0,
		},
		pb.Employee{
			BadgeNumber:         234,
			FirstName:           "Lisa",
			LastName:            "Wu",
			VacationAccrualRate: 1.7,
			VacationAccrued:     10,
			Vacations: []*pb.Vacation{
				&pb.Vacation{
					Id:        getNextVacationId(),
					StartDate: getLocalDate(2016, 5, 6).Unix(),
					Duration:  4,
				},
				&pb.Vacation{
					Id:        getNextVacationId(),
					StartDate: getLocalDate(2016, 5, 13).Unix(),
					Duration:  4,
				},
				&pb.Vacation{
					Id:        getNextVacationId(),
					StartDate: getLocalDate(2016, 5, 27).Unix(),
					Duration:  4,
				},
				&pb.Vacation{
					Id:        getNextVacationId(),
					StartDate: getLocalDate(2016, 5, 20).Unix(),
					Duration:  4,
				},
				&pb.Vacation{
					Id:        getNextVacationId(),
					StartDate: getLocalDate(2016, 6, 16).Unix(),
					Duration:  16,
				},
			},
		},
	}
	stream, err := client.SaveAll(context.Background())
	if err != nil {
		log.Fatal(err)
	}
	doneCh := make(chan struct{})
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				doneCh <- struct{}{}
				break
			}
			if err != nil {
				log.Fatal(err)
			}
			fmt.Println(res.Employee)
		}
	}()
	for _, e := range employees {
		err := stream.Send(&pb.EmployeeRequest{Employee: &e})
		if err != nil {
			log.Fatal(err)
		}
	}
	stream.CloseSend()
	<-doneCh
}

var (
	nextEmployeeId int = 0
	nextVacationId int = 0
)

func getNextVacationId() int32 {
	nextVacationId++
	return int32(nextVacationId)
}

func getNextEmployeeId() int {
	nextEmployeeId++
	return nextEmployeeId
}

func getLocalDate(year int, month time.Month, day int) *time.Time {
	loc, err := time.LoadLocation("America/New_York")
	if err != nil {
		log.Fatal(err)
	}
	result := time.Date(year, month, day, 0, 0, 0, 0, loc)
	return &result
}
