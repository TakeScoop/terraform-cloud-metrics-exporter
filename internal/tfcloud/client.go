package tfcloud

import "github.com/hashicorp/go-tfe"

type Client struct {
	client *tfe.Client
	config *tfe.Config
}

func New(config *tfe.Config) (*Client, error) {
	if config == nil {
		config = tfe.DefaultConfig()
	}

	client, err := tfe.NewClient(config)
	if err != nil {
		return nil, err
	}

	return &Client{
		client: client,
		config: config,
	}, nil
}
