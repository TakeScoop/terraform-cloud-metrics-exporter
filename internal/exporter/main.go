package exporter

import (
	"context"
	"log"
	"time"
)

type Config struct {
	Interval     time.Duration
	Organization string
}

type Exporter struct {
	interval     time.Duration
	organization string
}

func New(config *Config) *Exporter {
	return &Exporter{
		interval:     config.Interval,
		organization: config.Organization,
	}
}

func (e *Exporter) Start(ctx context.Context) error {
	ticker := time.NewTicker(e.interval)

	log.Printf("starting exporter with interval: %s", e.interval)

	for {
		select {
		case <-ctx.Done():
			log.Println("context canceled, stopping exporter")
			return ctx.Err()
		case <-ticker.C:
			log.Println("fetching agent data")
		}
	}
}
