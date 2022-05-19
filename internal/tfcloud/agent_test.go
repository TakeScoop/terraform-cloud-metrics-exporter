package tfcloud

import (
	"context"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/hashicorp/go-tfe"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestListAgents_ok(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v2/agent-pools/the-pool-id/agents", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)

		w.Write([]byte(`
			{
				"data": [
					{
						"attributes": {
							"name": "agent-1",
							"status": "running"
						}
					}
				]
			}
		`))
	})

	server := httptest.NewServer(mux)

	t.Cleanup(server.Close)

	config := tfe.DefaultConfig()
	config.Address = server.URL
	config.Token = "fake"

	client, err := New(config)
	require.NoError(t, err)

	agents, err := client.ListAgents(context.Background(), "the-pool-id")
	require.NoError(t, err)

	assert.Equal(t, []*Agent{
		{
			Name:   "agent-1",
			Status: "running",
		},
	}, agents)
}

func TestListAgents_error(t *testing.T) {
	mux := http.NewServeMux()

	mux.HandleFunc("/api/v2/agent-pools/the-pool-id/agents", func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, "GET", r.Method)
		w.WriteHeader(http.StatusForbidden)
	})

	server := httptest.NewServer(mux)

	t.Cleanup(server.Close)

	config := tfe.DefaultConfig()
	config.Address = server.URL
	config.Token = "fake"

	client, err := New(config)
	require.NoError(t, err)

	_, err = client.ListAgents(context.Background(), "the-pool-id")
	require.Error(t, err)
}
