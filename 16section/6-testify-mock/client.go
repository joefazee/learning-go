package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type HTTPClient interface {
	Get(url string) (*http.Response, error)
}

type PostsClient struct {
	httpClient HTTPClient
	baseURL    string
}

func NewPostsClient(httpClient HTTPClient, baseURL string) *PostsClient {
	return &PostsClient{
		httpClient: httpClient,
		baseURL:    baseURL,
	}
}

type Post struct {
	ID     int    `json:"id"`
	Title  string `json:"title"`
	Body   string `json:"body"`
	UserID int    `json:"userId"`
}

type PostsResponse struct {
	Posts []Post `json:"posts"`
	Total int    `json:"total"`
}

func (c *PostsClient) FetchPosts(limit int) ([]Post, error) {

	url := fmt.Sprintf("%s/posts?limit=%d", c.baseURL, limit)
	resp, err := c.httpClient.Get(url) // depend on interface
	if err != nil {
		return nil, fmt.Errorf("failed to fetch posts: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	var postsResp PostsResponse
	if err := json.NewDecoder(resp.Body).Decode(&postsResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}

	return postsResp.Posts, nil
}
