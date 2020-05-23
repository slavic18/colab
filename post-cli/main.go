package main

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"

	"context"

	// Import the generated protobuf code
	"github.com/micro/go-micro/v2"
	pb "github.com/slavic18/colab/post-service/proto/post"
)

const (
	address         = "localhost:50051"
	defaultFilename = "post.json"
)

func parseFile(file string) (*pb.Post, error) {
	var post *pb.Post
	data, err := ioutil.ReadFile(file)
	if err != nil {
		return nil, err
	}
	json.Unmarshal(data, &post)
	return post, err
}

func main() {
	service := micro.NewService(micro.Name("colab.service.cli"))
	service.Init()
	// Set up a connection to the server.
	client := pb.NewPostService("colab.service.post", service.Client())

	// Contact the server and print out its response.
	file := defaultFilename
	if len(os.Args) > 1 {
		file = os.Args[1]
	}

	post, err := parseFile(file)

	if err != nil {
		log.Fatalf("Could not parse file: %v", err)
	}

	r, err := client.CreatePost(context.Background(), post)

	if err != nil {
		log.Fatalf("Could not greet: %v", err)
	}
	log.Printf("Created: %t", r.Created)

	getPosts, err := client.GetPosts(context.Background(), &pb.GetRequest{})

	if err != nil {
		log.Fatalf("Could not list consignments: %v", err)
	}

	for _, v := range getPosts.Posts {
		log.Println(v)
	}
}
