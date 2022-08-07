package main

import (
	"context"
	"log"
	"net"

	pb "github.com/amogh2019/blog-service/blog/proto"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/reflection"
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

	// setting up tcp connection for server to use
	conn, err := net.Listen("tcp", serverAddress)
	if err != nil {
		log.Fatal("cannot listen to ", serverAddress, err)
	}
	defer conn.Close()

	// setting up ssl for server
	opts := []grpc.ServerOption{}
	tlsEnabled := true
	if tlsEnabled {
		cred, err := credentials.NewServerTLSFromFile("ssl/server.crt", "ssl/server.pem")
		if err != nil {
			log.Fatal("cannot load server cert", err)
		}
		opts = append(opts, grpc.Creds(cred))
	}

	// setting up server
	server := grpc.NewServer(opts...)
	pb.RegisterBlogServiceServer(server, &BlogServiceServerImpl{})

	reflection.Register(server)

	log.Println("Starting Blog Server at :", serverAddress)
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
