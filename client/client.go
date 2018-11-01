package client

// Config holds the login information for an API client
type Config struct {
	Email  string
	APIKey string
}

// Client holds one API client
type Client struct {
	config Config
}

// New creates a new API client with the config
func New(conf Config) (c Client) {
	c.config = conf
	return
}
