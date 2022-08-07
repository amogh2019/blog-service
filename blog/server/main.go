package main

import (
	"context"
	"log"
	"net"

	pb "github.com/amogh2019/blog-service/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
)

var mongoURI = "mongodb://root:root@localhost:27017/"
var blogCollection *mongo.Collection

const serverAddress = "0.0.0.0:50052"

type BlogServiceServerImpl struct {
	pb.BlogServiceServer
}

func main() {
	// setting mongo connection // getting pointer to collection
	blogCollection = connectToMongoAndGetCollection(mongoURI, "blogdb", "blog")
	// setting up server
	conn, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatal("cannot listen to ", serverAddress, err)
	}
	defer conn.Close()

	server := grpc.NewServer()
	pb.RegisterBlogServiceServer(server, &BlogServiceServerImpl{})

	err = server.Serve(conn)
	if err != nil {
		log.Fatal("cannot start server", err)
	}
}

func connectToMongoAndGetCollection(mongoURI string, dbname string, collectionName string) *mongo.Collection {
	mongoClient, err := mongo.NewClient(options.Client().ApplyURI(mongoURI))
	if err != nil {
		log.Fatal("error in setting up a connection to mongo", err)
	}
	err = mongoClient.Connect(context.Background())
	if err != nil {
		log.Fatal("error in connecting to mongo", err)
	}
	return mongoClient.Database(dbname).Collection(collectionName)
}
