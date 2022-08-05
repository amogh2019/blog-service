all: blog

blog: buildBlogProto

buildBlogProto : 
	protoc --go_opt=module=github.com/amogh2019/blog-service --go_out=. --go-grpc_opt=module=github.com/amogh2019/blog-service --go-grpc_out=.   blog/proto/*.proto
