package fastbill

import "github.com/fastbill/sdk-go/customer"

// ClientConfig holds the login information for an API client
type ClientConfig struct {
	Email  string
	APIKey string
}

// Client holds one API client
type Client struct {
	config   *ClientConfig
	Customer *customer.Customer
}

// NewClient creates a new API client from the config and returns it
func NewClient(conf *ClientConfig) (cl Client, err error) {
	// TODO validate configuration
	cl.config = conf
	return
}
