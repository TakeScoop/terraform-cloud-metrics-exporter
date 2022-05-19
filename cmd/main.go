package cmd

import (
	"context"
	"errors"
	"flag"
	"os/signal"
	"time"

	"github.com/takescoop/terraform-cloud-metrics-exporter/internal/exporter"
)

func Exec(args []string) error {
	ctx := context.Background()

	ctx, cancel := signal.NotifyContext(ctx)
	defer cancel()

	flags := flag.NewFlagSet("terraform-cloud-metrics-exporter", flag.ExitOnError)

	var organization string
	flags.StringVar(&organization, "organization", "", "The Terraform Cloud organization")

	var interval time.Duration
	flags.DurationVar(&interval, "interval", time.Minute, "The interval to fetch agent data")

	if err := flags.Parse(args); err != nil {
		return err
	}

	if organization == "" {
		return errors.New("organization is required")
	}

	return exporter.New(&exporter.Config{
		Organization: organization,
		Interval:     interval,
	}).Start(ctx)
}
