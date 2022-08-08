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
)

func (server *BlogServiceServerImpl) ReadBlog(ctx context.Context, blogIdDto *pb.BlogIdDto) (*pb.BlogDto, error) {

	log.Println("ReadBlog invoked : ", blogIdDto)

	// string to ObjectId
	objectId, err := primitive.ObjectIDFromHex(blogIdDto.BlogId)

	if err != nil {
		errString := fmt.Sprintf("cannot parse into objectid %v\n %v", blogIdDto.BlogId, err)
		log.Println(errString)
		return nil, status.Errorf(
			codes.InvalidArgument,
			errString,
		)
	}

	// finding blog
	responseData := &BlogItem{}
	query := bson.M{
		"_id": objectId,
	}
	result := blogCollection.FindOne(ctx, query)

	if err := result.Decode(responseData); err != nil {
		errString := fmt.Sprintf("No Blog found for id %v\n %v", blogIdDto.BlogId, err)
		log.Println(errString)
		return nil, status.Errorf(
			codes.NotFound,
			errString,
		)
	}

	// returning response
	blogDto := documentToBlogDto(responseData)
	log.Println("Blog", blogDto)
	return blogDto, nil

}
