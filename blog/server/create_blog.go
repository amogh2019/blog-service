package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/amogh2019/blog-service/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *BlogServiceServerImpl) CreateBlog(ctx context.Context, requestDto *pb.BlogDto) (*pb.BlogIdDto, error) {

	log.Println("CreateBlog invoked : ", requestDto)

	// inserting
	insertOneResult, err := blogCollection.InsertOne(ctx, *blogDtoToDocumentWithoutId(requestDto))
	if err != nil {
		log.Println(err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while inserting %v\n %v", requestDto, err),
		)
	}
	// checking inserted id
	oid, ok := insertOneResult.InsertedID.(primitive.ObjectID) // parsing interface specifically into objectid
	if !ok {
		log.Println(err)
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error from while inserting %v\n %v", requestDto, err),
		)
	}

	// returning response
	log.Println("CreateBlog complete : blogId ", oid)
	return &pb.BlogIdDto{
		BlogId: oid.Hex(),
	}, nil

}
