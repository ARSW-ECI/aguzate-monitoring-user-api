package main

import (
	"context"
	"eci/aguzate_monitoring/user_api/src/eci/aguzate_monitoring"
	"flag"
	"fmt"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"io"
	"log"
)

const port = ":9000"

func main() {
	option := flag.Int("o", 1, "Command to Run")
	flag.Parse()

	fmt.Printf("Option is %v\n", *option)

	opts := []grpc.DialOption{grpc.WithTransportCredentials(insecure.NewCredentials())}
	//conn, err := grpc.Dial("localhost"+port, opts...)
	conn, err := grpc.Dial("localhost"+port, opts...)

	if err != nil {
		log.Fatal(err)
	}

	defer conn.Close()
	client := aguzate_monitoring.NewBikeServiceClient(conn)

	switch *option {
	case 1:
		GetBikesByUserId(client)
	case 2:
		UpdateBike(client)
	case 3:
		GetBikeLocation(client)
	}
}

func GetBikesByUserId(client aguzate_monitoring.BikeServiceClient) {
	stream, err := client.GetBikesByUserId(context.Background(), &aguzate_monitoring.GetBikesByUserIdRequest{UserId: "62b1cb32d5d85b53a5783637"})

	if err != nil {
		log.Fatal(err)
	}

	for {
		b, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(b.Bike)
	}
}

func UpdateBike(client aguzate_monitoring.BikeServiceClient) {
	var bike = aguzate_monitoring.Bike{
		BikeId: "1",
		Color:  "red",
		Size:   "L",
	}

	b, err := client.UpdateBike(context.Background(), &aguzate_monitoring.UpdateBikeRequest{Bike: &bike})

	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(b)
}

func GetBikeLocation(client aguzate_monitoring.BikeServiceClient) {
	stream, err := client.GetBikeLocation(context.Background(), &aguzate_monitoring.GetBikeLocationRequest{BikeId: "1", UserId: "1"})

	if err != nil {
		log.Fatal(err)
	}

	for {
		l, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println(l.Location)
	}
}
