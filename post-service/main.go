package main

import (
	"fmt"

	"golang.org/x/net/context"

	// Import the generated protobuf code
	pb "github.com/slavic18/colab/post-service/proto/post"

	"github.com/micro/go-micro/v2"
)

const (
	port = ":50051"
)

type repository interface {
	Create(*pb.Post) (*pb.Post, error)
	GetAll() []*pb.Post
}

// Repository - Dummy repository, this simulates the use of a datastore
// of some kind. We'll replace this with a real implementation later on.
type Repository struct {
	posts []*pb.Post
}

// Create a new post
func (repo *Repository) Create(post *pb.Post) (*pb.Post, error) {
	updated := append(repo.posts, post)
	repo.posts = updated
	return post, nil
}

// GetAll posts
func (repo *Repository) GetAll() []*pb.Post {
	return repo.posts
}

// Service should implement all of the methods to satisfy the service
// we defined in our protobuf definition. You can check the interface
// in the generated code itself for the exact method signatures etc
// to give you a better idea.
type service struct {
	repo repository
}

// CreatePost - we created just one method on our service,
// which is a create method, which takes a context and a request as an
// argument, these are handled by the gRPC server.
func (s *service) CreatePost(ctx context.Context, req *pb.Post, res *pb.Response) error {

	// Save our post
	post, err := s.repo.Create(req)
	if err != nil {
		return err
	}

	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Created = true
	res.Post = post

	return nil
}

func (s *service) GetPosts(ctx context.Context, req *pb.GetRequest, res *pb.Response) error {

	posts := s.repo.GetAll()
	// Return matching the `Response` message we created in our
	// protobuf definition.
	res.Posts = posts

	return nil

}

func main() {

	repo := &Repository{}

	// create a new service. Optionally include some options here
	srv := micro.NewService(
		// this name must match the package name given in your protobuf definition
		micro.Name("colab.service.post"),
	)

	// Init will parse the command line flags.
	srv.Init()

	pb.RegisterPostServiceHandler(srv.Server(), &service{repo})

	if err := srv.Run(); err != nil {
		fmt.Println(err)
	}

}
