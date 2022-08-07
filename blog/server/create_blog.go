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

func CreateBlog(ctx context.Context, requestDto *pb.BlogDto) (*pb.BlogIdDto, error) {

	log.Println("CreateBlog invoked : ", requestDto)

	// inserting
	insertOneResult, err := blogCollection.InsertOne(ctx, blogDtoToDocumentWithoutId(requestDto))
	if err != nil {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error while inserting %v\n", requestDto),
		)
	}
	// checking inserted id
	oid, ok := insertOneResult.InsertedID.(primitive.ObjectID)
	if !ok {
		return nil, status.Errorf(
			codes.Internal,
			fmt.Sprintf("Error from while inserting %v\n", requestDto),
		)
	}

	// returning response
	return &pb.BlogIdDto{
		BlogId: oid.Hex(),
	}, nil

}
