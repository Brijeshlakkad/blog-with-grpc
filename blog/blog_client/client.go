package main

import (
	"context"
	"fmt"
	"grpc-go-example/blog/blogpb"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm a client")

	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())

	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	defer cc.Close()

	c := blogpb.NewBlogServiceClient(cc)

	// Create blog
	// fmt.Println("Creating the blog")
	// blog := blogpb.Blog{
	// 	AuthorId: "Stephane",
	// 	Title:    "My First Blog",
	// 	Content:  "Content of the first blog",
	// }
	// createBlogRes, err := c.CreateBlog(context.Background(), &blogpb.CreateBlogRequest{Blog: &blog})
	// if err != nil {
	// 	log.Fatalf("Unexpected error: %v", err)
	// }
	// fmt.Printf("Blog has been created: %v", createBlogRes)

	// Read blog
	// 1. Wrong blog id.
	// fmt.Println("Reading the blog")
	// res, err := c.ReadBlog(context.Background(), &blogpb.ReadBlogRequest{
	// 	BlogId: "111",
	// })
	// if err != nil {
	// 	log.Fatalf("Error happened while reading: %v", err)
	// }
	// fmt.Printf("Blog was read %v", res)

	// 2. Correct blog id.
	// readBlogReq := &blogpb.ReadBlogRequest{BlogId: "61cfd7aa1a9d34a8c2ea3419"}
	// readBlogRes, err := c.ReadBlog(context.Background(), readBlogReq)
	// if err != nil {
	// 	fmt.Sprintf("Error happened while reading: %v", err)
	// }
	// fmt.Printf("Blog was read %v", readBlogRes)

	// Update blog
	// newBlog := &blogpb.Blog{
	// 	Id:       "61cfd7aa1a9d34a8c2ea3419",
	// 	AuthorId: "Changed Author",
	// 	Title:    "My First Blog (edited)",
	// 	Content:  "Content of the first blog, with some awesome additions!",
	// }

	// updateRes, updateErr := c.UpdateBlog(context.Background(), &blogpb.UpdateBlogRequest{
	// 	Blog: newBlog,
	// })
	// if updateErr != nil {
	// 	fmt.Printf("Error happened while updating: %v \n", updateErr)
	// }
	// fmt.Printf("Blog was read: %v", updateRes)

	// Delete blog
	// deleteRes, deleteErr := c.DeleteBlog(context.Background(), &blogpb.DeleteBlogRequest{
	// 	BlogId: "61cfd7aa1a9d34a8c2ea3419",
	// })

	// if deleteErr != nil {
	// 	fmt.Printf("Error happened while deleting %v \n", deleteErr)
	// }
	// fmt.Printf("Blog was deleted: %v", deleteRes)

	// List blogs
	resStream, err := c.ListBlog(context.Background(), &blogpb.ListBlogRequest{})
	if err != nil {
		log.Fatalf("error while calling ListBlog RPC: %v", err)
	}
	for {
		res, err := resStream.Recv()
		if err == io.EOF {
			// we've reached the end of the stream
			break
		}
		if err != nil {
			log.Fatalf("error while reading stream: %v", err)
		}
		fmt.Println(res.GetBlog())
	}
}
