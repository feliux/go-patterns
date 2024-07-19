package main

import (
	"crypto/tls"
	"net/http"
	"time"
)

// Client includes parameter for the HTTP client.
type Client struct {
	client          *http.Client
	timeout         time.Duration
	userAgent       string
	followRedirects bool
}

// Option is a functional option type that allows us to configure the Client.
type Option func(*Client)

// NewClient creates the HTTP client with custom options.
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

// WithTimeout set the timeout for the client request.
func WithTimeout(timeout time.Duration) Option {
	return func(c *Client) {
		c.timeout = timeout
	}
}

// WithUserAgent the user-agent of the client.
func WithUserAgent(userAgent string) Option {
	return func(c *Client) {
		c.userAgent = userAgent
	}
}

// WithRedirects configue the client to follow redirects.
func WithoutRedirects() Option {
	return func(c *Client) {
		c.followRedirects = false
	}
}

// UseInsecureTransport configure the client to no verify trusted certificates.
func UseInsecureTransport() Option {
	return func(c *Client) {
		c.client.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
	}
}

// Get performs a HTTP GET passing the custom client object.
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
