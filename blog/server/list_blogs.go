package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/amogh2019/blog-service/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func (server *BlogServiceServerImpl) ListBlogs(req *emptypb.Empty, stream pb.BlogService_ListBlogsServer) error {

	log.Println("List Blogs invoked")

	// updating blog
	cursor, err := blogCollection.Find(context.Background(), primitive.D{{}})

	if err != nil {
		errString := fmt.Sprintf("mongo error finding all documents\n %v", err)
		log.Println(errString)
		return status.Errorf(
			codes.Internal,
			errString,
		)
	}
	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()) {
		data := &BlogItem{}
		err := cursor.Decode(data)

		if err != nil {
			errString := fmt.Sprintf("error in decoding cursor data into blogitem entity \n %v", err)
			log.Println(errString)
			return status.Errorf(
				codes.Internal,
				errString,
			)
		}

		blogDto := documentToBlogDto(data)
		log.Println("Blog", blogDto)
		stream.Send(documentToBlogDto(data))
	}

	if err := cursor.Err(); err != nil {
		errString := fmt.Sprintf("error in cursor while list blog \n %v", err)
		log.Println(errString)
		return status.Errorf(
			codes.Internal,
			errString,
		)
	}

	// returning response
	log.Println("List Blogs complete")
	return nil

}
