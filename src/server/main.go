package main

import (
	"context"
	"eci/aguzate_monitoring/user_api/src/eci/aguzate_monitoring"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const port = ":9000"

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	aguzate_monitoring.RegisterBikeServiceServer(s, new(bikeService))
	log.Println("Starting server on port " + port)
	s.Serve(lis)
}

type bikeService struct{}

func (s *bikeService) GetBikesByUserId(
	req *aguzate_monitoring.GetBikesByUserIdRequest,
	stream aguzate_monitoring.BikeService_GetBikesByUserIdServer) error {
	var bikes = []aguzate_monitoring.Bike{
		aguzate_monitoring.Bike{
			BikeId: "1",
			Color:  "red",
			Size:   "L",
		},
	}

	for _, b := range bikes {
		stream.Send(&aguzate_monitoring.BikeResponse{Bike: &b})
	}

	return nil

	return nil
}

func (s *bikeService) UpdateBike(
	context.Context,
	*aguzate_monitoring.UpdateBikeRequest) (
	*aguzate_monitoring.BikeResponse,
	error) {
	var bike = aguzate_monitoring.Bike{
		BikeId: "1",
		Color:  "red",
		Size:   "L",
	}
	return &aguzate_monitoring.BikeResponse{Bike: &bike}, nil
}

func (s *bikeService) GetBikeLocation(
	req *aguzate_monitoring.GetBikeLocationRequest,
	stream aguzate_monitoring.BikeService_GetBikeLocationServer) error {

	var location = aguzate_monitoring.Location{
		Latitude:  123423,
		Longitude: 234345,
	}
	var bikeLocation = aguzate_monitoring.BikeLocation{
		BikeId:   "1",
		Location: &location,
	}

	for i := 1; i < 10; i++ {
		stream.Send(&aguzate_monitoring.BikeLocationResponse{
			Location: &bikeLocation,
		})
		time.Sleep(5 * time.Second)
	}

	return nil
}
