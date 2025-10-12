package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {

	baseURL := "https://dummyjson.com"

	postsClient := NewPostsClient(http.DefaultClient, baseURL)

	posts, err := postsClient.FetchPosts(3)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}

	for _, post := range posts {
		fmt.Printf("[%d] %s\n", post.ID, post.Title)
	}

}
