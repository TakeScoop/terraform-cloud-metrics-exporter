package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

type Agent struct {
	Name   string `json:"name"`
	Status string `json:"status"`
}

func ListAgents(pool string) ([]*Agent, error) {
	url := fmt.Sprintf("https://app.terraform.io/api/v2/agent-pools/%s/agents", pool)

	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Authorization", fmt.Sprintf("Bearer %s", os.Getenv("TFE_TOKEN")))

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to list agents: %d", resp.StatusCode)
	}

	var result AgentsResponse

	defer resp.Body.Close()

	if err = json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return nil, err
	}

	return result.Agents(), nil
}

type AgentsResponse struct {
	Data []struct {
		Attributes *Agent `json:"attributes"`
	} `json:"data"`
}

func (r *AgentsResponse) Agents() []*Agent {
	agents := make([]*Agent, len(r.Data))

	for i, d := range r.Data {
		agents[i] = d.Attributes
	}

	return agents
}
