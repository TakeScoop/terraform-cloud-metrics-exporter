package agentstatus

import (
	"context"

	"github.com/takescoop/terraform-cloud-metrics-exporter/internal/tfcloud"
)

// AgentPool represents a Terraform Cloud agent pool, including its agents.
type AgentPool struct {
	*tfcloud.AgentPool
	Agents []*tfcloud.Agent
}

// ByStatus returns the number of agents in each status for the pool.
func (a *AgentPool) ByStatus() map[string]uint {
	status := make(map[string]uint)

	for _, agent := range a.Agents {
		status[agent.Status]++
	}

	return status
}

// Summary is a summary of the current status of agent pools and their agents.
type Summary struct {
	Pools []*AgentPool
}

// Get returns a summary for the specified organization.
func Get(ctx context.Context, client *tfcloud.Client, organization string) (*Summary, error) {
	pools, err := client.ListAgentPools(ctx, organization)
	if err != nil {
		return nil, err
	}

	summary := &Summary{}
	for _, pool := range pools {
		agents, err := client.ListAgents(ctx, pool.ID)
		if err != nil {
			return nil, err
		}

		summary.Pools = append(summary.Pools, &AgentPool{
			AgentPool: pool,
			Agents:    agents,
		})
	}

	return summary, nil
}
