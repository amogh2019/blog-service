package main

import (
	"bufio"
	"context"
	"log"
	"strings"

	pb "github.com/amogh2019/blog-service/blog/proto"
	"google.golang.org/grpc/status"
)

func ReadBlogById(blogServiceClient *pb.BlogServiceClient, scanner *bufio.Scanner) {

	log.Println("Preparing to get blogId to read from db")

	log.Println("Enter BlogId")
	(*scanner).Scan()
	blogIdString := (*scanner).Text()
	blogIdString = strings.TrimSpace(blogIdString)

	blogDto, err := (*blogServiceClient).ReadBlog(context.Background(), &pb.BlogIdDto{BlogId: blogIdString})
	if err != nil {
		grpcParsedStatus, ok := status.FromError(err)
		if ok {
			log.Println("ErrorCode : ", grpcParsedStatus.Code(), " ErrorMsg: ", grpcParsedStatus.Message())
		} else {
			log.Fatal("Non-gRPC Error from server", err)
		}
		return
	}

	log.Println("Blog Found : ", blogDto)
}
