package main

import (
	"context"

	"go.opentelemetry.io/otel/exporters/stdout/stdoutmetric"
	"go.opentelemetry.io/otel/metric/global"
	"go.opentelemetry.io/otel/metric/instrument"
	controller "go.opentelemetry.io/otel/sdk/metric/controller/basic"
	processor "go.opentelemetry.io/otel/sdk/metric/processor/basic"
	"go.opentelemetry.io/otel/sdk/metric/selector/simple"
)

func main() {
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

	gauge.Observe(context.TODO(), int64(1))

	if err = pusher.Stop(context.TODO()); err != nil {
		panic(err)
	}
}
