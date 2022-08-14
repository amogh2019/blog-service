package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/amogh2019/blog-service/blog/proto"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *BlogServiceServerImpl) DeleteBlog(ctx context.Context, idDto *pb.BlogIdDto) (*emptypb.Empty, error) {

	log.Println("Delete invoked : ", idDto)

	// string to ObjectId
	objectId, err := primitive.ObjectIDFromHex(idDto.BlogId)

	if err != nil {
		errString := fmt.Sprintf("cannot parse into objectid %v\n %v", idDto.BlogId, err)
		log.Println(errString)
		return nil, status.Errorf(
			codes.InvalidArgument,
			errString,
		)
	}

	// sending find and delete on mongo
	deleteResult, err := blogCollection.DeleteOne(ctx,
		bson.M{
			"_id": objectId,
		},
	)

	if err != nil {
		errString := fmt.Sprintf("mongo error in deleting document\n %v", err)
		log.Println(errString)
		return nil, status.Errorf(
			codes.Internal,
			errString,
		)
	}

	if deleteResult.DeletedCount < 1 {
		errString := fmt.Sprintf("no document matched for selected filter(search by id). %v \n %v", idDto, err)
		log.Println(errString)
		return nil, status.Errorf(
			codes.Internal,
			errString,
		)
	}

	// returning response
	log.Println("Deleted", idDto)
	return &emptypb.Empty{}, nil
}
