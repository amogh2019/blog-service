package main

import (
	"bufio"
	"context"
	"log"

	pb "github.com/amogh2019/blog-service/blog/proto"
	"google.golang.org/grpc/status"
)

func CreateBlog(blogServiceClient *pb.BlogServiceClient, scanner *bufio.Scanner) {

	log.Println("Preparing Create Blog Payload")

	createBlogDto := pb.BlogDto{}
	log.Println("Enter Blog Title")
	(*scanner).Scan()
	createBlogDto.Title = (*scanner).Text()
	log.Println("Enter Blog Content")
	(*scanner).Scan()
	createBlogDto.Content = (*scanner).Text()
	log.Println("Enter Blog Author")
	(*scanner).Scan()
	createBlogDto.AuthorId = (*scanner).Text()

	blogIdDto, err := (*blogServiceClient).CreateBlog(context.Background(), &createBlogDto)
	if err != nil {
		grpcParsedStatus, ok := status.FromError(err)
		if ok {
			log.Println("ErrorCode : ", grpcParsedStatus.Code(), " ErrorMsg: ", grpcParsedStatus.Message())
		} else {
			log.Fatal("Non-gRPC Error from server", err)
		}
		return
	}

	createBlogDto.Id = blogIdDto.BlogId
	log.Println("Blog Created. BlogId : ", blogIdDto.BlogId)
}
