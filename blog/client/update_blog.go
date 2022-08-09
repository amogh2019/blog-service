package main

import (
	"bufio"
	"context"
	"log"
	"strings"

	pb "github.com/amogh2019/blog-service/blog/proto"
	"google.golang.org/grpc/status"
)

func UpdateBlogById(blogServiceClient *pb.BlogServiceClient, scanner *bufio.Scanner) {

	log.Println("Preparing to get blogId to read from db")

	createBlogDto := pb.BlogDto{}
	log.Println("Enter BlogId")
	(*scanner).Scan()
	blogIdString := (*scanner).Text()
	createBlogDto.Id = strings.TrimSpace(blogIdString)

	log.Println("Enter Blog Title")
	(*scanner).Scan()
	createBlogDto.Title = (*scanner).Text()
	log.Println("Enter Blog Content")
	(*scanner).Scan()
	createBlogDto.Content = (*scanner).Text()
	log.Println("Enter Blog Author")
	(*scanner).Scan()
	createBlogDto.AuthorId = (*scanner).Text()

	_, err := (*blogServiceClient).UpdateBlog(context.Background(), &createBlogDto)
	if err != nil {
		grpcParsedStatus, ok := status.FromError(err)
		if ok {
			log.Println("ErrorCode : ", grpcParsedStatus.Code(), " ErrorMsg: ", grpcParsedStatus.Message())
		} else {
			log.Fatal("Non-gRPC Error in from server", err)
		}
		return
	}

	log.Println("Blog Updated ")
}
