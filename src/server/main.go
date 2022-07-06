package main

import (
	"context"
	"eci/aguzate_monitoring/user_api/src/eci/aguzate_monitoring"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"log"
	"net"
	"time"
)

const port = ":9000"

var mongoClient *mongo.Client

func main() {
	mongoClient = MongoConnect()
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal(err)
	}
	s := grpc.NewServer()
	aguzate_monitoring.RegisterBikeServiceServer(s, new(bikeService))
	log.Println("Starting server on port " + port)
	s.Serve(lis)
}

func MongoConnect() *mongo.Client {
	clientOptions := options.Client().ApplyURI("mongodb+srv://aguzate-admin:aguzate@cluster0.vte0q.mongodb.net/aguzate")
	mongoClient, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to Mongo...")
	return mongoClient
}

type Bike struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Color string             `bson:"color"`
	Size  string             `bson:"size"`
}

type User struct {
	ID    primitive.ObjectID `bson:"_id,omitempty"`
	Bikes []Bike             `bson:"bikes"`
}

type bikeService struct{}

func (s *bikeService) GetBikesByUserId(
	req *aguzate_monitoring.GetBikesByUserIdRequest,
	stream aguzate_monitoring.BikeService_GetBikesByUserIdServer) error {

	user := User{}
	docId, _ := primitive.ObjectIDFromHex(req.UserId)
	collection := mongoClient.Database("aguzate").Collection("users")
	collection.FindOne(context.TODO(), bson.M{"_id": docId}).Decode(&user)

	for _, b := range user.Bikes {
		stream.Send(&aguzate_monitoring.BikeResponse{Bike: &aguzate_monitoring.Bike{
			BikeId: b.ID.String(),
			Color:  b.Color,
			Size:   b.Size,
		}})
	}

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
