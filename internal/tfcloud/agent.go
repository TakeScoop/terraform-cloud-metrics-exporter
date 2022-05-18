package tfcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// ListAgents lists the agents in a given pool
// TODO: use go-tfe when support for this API is added
func (c *Client) ListAgents(ctx context.Context, poolId string) ([]*Agent, error) {
	url := fmt.Sprintf("%s%sagent-pools/%s/agents", c.config.Address, c.config.BasePath, poolId)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req = req.WithContext(ctx)
	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", c.config.Token))

	resp, err := c.config.HTTPClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list agents: %d", resp.StatusCode)
	}

	var result agentsResponse
	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Agents(), nil
}

type Agent struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

type agentsResponse struct {
	Data []struct {
		Attributes *Agent `json:"attributes"`
	} `json:"data"`
}

func (r *agentsResponse) Agents() []*Agent {
	agents := make([]*Agent, len(r.Data))

	for i, d := range r.Data {
		agents[i] = d.Attributes
	}

	return agents
}
