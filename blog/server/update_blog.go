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

func (server *BlogServiceServerImpl) UpdateBlog(ctx context.Context, reqDto *pb.BlogDto) (*emptypb.Empty, error) {

	log.Println("Update invoked : ", reqDto)

	// string to ObjectId
	objectId, err := primitive.ObjectIDFromHex(reqDto.Id)

	if err != nil {
		errString := fmt.Sprintf("cannot parse into objectid %v\n %v", reqDto, err)
		log.Println(errString)
		return nil, status.Errorf(
			codes.InvalidArgument,
			errString,
		)
	}

	// updating blog
	res, err := blogCollection.UpdateOne(
		ctx,
		bson.M{
			"_id": objectId,
		},
		bson.M{
			"$set": blogDtoToDocumentWithoutId(reqDto),
		},
	)

	if err != nil {
		errString := fmt.Sprintf("mongo error in updating document\n %v", err)
		log.Println(errString)
		return nil, status.Errorf(
			codes.Internal,
			errString,
		)
	}

	if res.MatchedCount < 1 {
		errString := fmt.Sprintf("no document matched for selected filter(search by id). %v \n %v", reqDto, err)
		log.Println(errString)
		return nil, status.Errorf(
			codes.Internal,
			errString,
		)
	}

	// returning response
	log.Println("Updated", reqDto)
	return &emptypb.Empty{}, nil

}
