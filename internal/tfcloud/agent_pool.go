package tfcloud

import (
	"context"

	"github.com/hashicorp/go-tfe"
)

type AgentPool = tfe.AgentPool

func (c *Client) ListAgentPools(ctx context.Context, organization string) ([]*AgentPool, error) {
	pools, err := c.client.AgentPools.List(ctx, organization, &tfe.AgentPoolListOptions{})
	if err != nil {
		return nil, err
	}

	return pools.Items, nil
}
