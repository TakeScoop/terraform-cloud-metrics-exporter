package main

import (
	"context"

	"github.com/takescoop/terraform-cloud-metrics-exporter/internal/agentstatus"
	"github.com/takescoop/terraform-cloud-metrics-exporter/internal/tfcloud"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
)

func main() {
	client, err := tfcloud.New(nil)
	if err != nil {
		panic(err)
	}

	exporter, err := stdoutmetric.New(stdoutmetric.WithPrettyPrint())
	if err != nil {
		panic(err)
	}

	pusher := controller.New(
		processor.NewFactory(
			simple.NewWithInexpensiveDistribution(),
			exporter,
		),
		controller.WithExporter(exporter),
	)

	if err = pusher.Start(context.TODO()); err != nil {
		panic(err)
	}

	global.SetMeterProvider(pusher)
	meter := global.Meter("terraform_cloud")

	gauge, err := meter.AsyncInt64().Gauge(
		"agents",
		instrument.WithDescription("The count of Terraform Cloud agents"),
	)

	if err != nil {
		panic(err)
	}

	summary, err := agentstatus.Get(context.TODO(), client, "takescoop")
	if err != nil {
		panic(err)
	}

	for _, pool := range summary.Pools {
		for status, count := range pool.ByStatus() {
			gauge.Observe(
				context.TODO(),
				int64(count),
				attribute.String("pool", pool.Name),
				attribute.String("status", status),
			)
		}
	}

	if err = pusher.Stop(context.TODO()); err != nil {
		panic(err)
	}
}
