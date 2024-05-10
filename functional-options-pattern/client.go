package main

import (
	"crypto/tls"
	"net/http"
	"time"
)

type Client struct {
	client          *http.Client
	timeout         time.Duration
	userAgent       string
	followRedirects bool
}

// Option is a functional option type that allows us to configure the Client.
type Option func(*Client)

func NewClient(options ...Option) *Client {
	client := &Client{
		client:          &http.Client{},
		timeout:         30 * time.Second,
		userAgent:       "My HTTP Client",
		followRedirects: true,
	}
	// Apply all the functional options to configure the client.
	for _, opt := range options {
		opt(client)
	}
	return client
}

func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.timeout = timeout
	}
}

func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

func WithoutRedirects() Option {
	return func(c *Client) {
		c.followRedirects = false
	}
}

func UseInsecureTransport() Option {
	return func(c *Client) {
		c.client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
}

func (c *Client) Get(url string) (*http.Response, error) {
	// Use c.client with all the configured options to perform the request.
	// ...
}

func main() {
	client := NewClient(
		WithTimeout(10*time.Second),
		WithUserAgent("My Custom User Agent"),
		UseInsecureTransport(),
	)
	response, err := client.Get("https://api.example.com/data")
	if err != nil {
		// Handle the error
	}
	defer response.Body.Close()
	// Process the response
}
