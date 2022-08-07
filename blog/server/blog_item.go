package main

import (
	pb "github.com/amogh2019/blog-service/blog/proto"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type BlogItem struct {
	ID       primitive.ObjectID `bson:"_id, omitempty"`
	AuthorId string             `bson:"author_id"`
	Title    string             `bson:"title"`
	Content  string             `bson:"content"`
}

func documentToBlogDto(data *BlogItem) *pb.BlogDto {
	return &pb.BlogDto{
		Id:       data.ID.Hex(),
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}
}

func blogDtoToDocumentWithoutId(data *pb.BlogDto) *BlogItem {
	return &BlogItem{
		AuthorId: data.AuthorId,
		Title:    data.Title,
		Content:  data.Content,
	}
}
