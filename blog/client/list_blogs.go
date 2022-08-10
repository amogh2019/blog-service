package main

import (
	"context"
	"io"
	"log"

	pb "github.com/amogh2019/blog-service/blog/proto"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/emptypb"
)

func FindAllBlogs(blogServiceClient *pb.BlogServiceClient) {

	log.Println("Preparing to read all blogs from db")

	stream, err := (*blogServiceClient).ListBlogs(context.Background(), &emptypb.Empty{})
	if err != nil {
		grpcParsedStatus, ok := status.FromError(err)
		if ok {
			log.Println("ErrorCode : ", grpcParsedStatus.Code(), " ErrorMsg: ", grpcParsedStatus.Message())
		} else {
			log.Fatal("Non-gRPC Error in from server", err)
		}
		return
	}

	i := 0
	for {
		dto, err := stream.Recv()

		if err == io.EOF {
			break
		}

		if err != nil {
			log.Fatal("Error while reading server response stream", err)
			return
		}

		i++
		log.Println("Data ", i, " \n", dto)

	}

	log.Println("List Blog complete")
}
