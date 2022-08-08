package main

import (
	"bufio"
	"log"
	"os"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/credentials/insecure"

	pb "github.com/amogh2019/blog-service/blog/proto"
)

var serverAddress string = "localhost:50052"

func main() {

	// setting up ssl for client
	tlsEnabled := true
	opts := []grpc.DialOption{}
	if tlsEnabled {
		cred, err := credentials.NewClientTLSFromFile("ssl/ca.crt", "")
		if err != nil {
			log.Fatal("error in loading CA trust certificate", err)
		}
		opts = append(opts, grpc.WithTransportCredentials(cred))
	} else {
		opts = append(opts, grpc.WithTransportCredentials(insecure.NewCredentials()))
	}

	// setting up connection to server
	conn, err := grpc.Dial(serverAddress, opts...)
	if err != nil {
		log.Fatal("error in starting dialup", serverAddress, err) // TODO this is not breaking when server is down // check how to write correct code to block this when server is not up
	}
	defer conn.Close()

	client := pb.NewBlogServiceClient(conn)

	scanner := bufio.NewScanner(os.Stdin)
	askForPress()
	shouldContinue := true
	for shouldContinue && scanner.Scan() {
		text := scanner.Text()
		switch text {
		case "1":
			CreateBlog(&client, scanner)
		case "2":
			ReadBlogById(&client, scanner)
		case "exit":
			log.Println("closing blog client")
			shouldContinue = false
		default:
			log.Println("Dhang se daalo!")
		}
		if shouldContinue {
			askForPress()
		}
	}

}

func askForPress() {
	log.Println("Please type the label for its action \n1 :  Create Blog \n2 :  Read Blog \nexit : To close Blog Client")
}
