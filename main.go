package main

import (
	"log"
	"os"

	"github.com/takescoop/terraform-cloud-metrics-exporter/cmd"
)

func main() {
	if err := cmd.Exec(os.Args[1:]); err != nil {
		log.Fatal(err)
	}

	// ctx := context.Background()

	// client, err := tfcloud.New(nil)
	// if err != nil {
	// 	panic(err)
	// }

	// exporter, err := stdoutmetric.New(stdoutmetric.WithPrettyPrint())
	// if err != nil {
	// 	panic(err)
	// }

	// pusher := controller.New(
	// 	processor.NewFactory(
	// 		simple.NewWithInexpensiveDistribution(),
	// 		exporter,
	// 	),
	// 	controller.WithExporter(exporter),
	// )

	// if err = pusher.Start(ctx); err != nil {
	// 	panic(err)
	// }

	// global.SetMeterProvider(pusher)
	// meter := global.Meter("terraform_cloud")

	// gauge, err := meter.AsyncInt64().Gauge(
	// 	"agents",
	// 	instrument.WithDescription("The count of Terraform Cloud agents"),
	// )

	// if err != nil {
	// 	panic(err)
	// }

	// summary, err := agentstatus.Get(ctx, client, "takescoop")
	// if err != nil {
	// 	panic(err)
	// }

	// for _, pool := range summary.Pools {
	// 	for status, count := range pool.ByStatus() {
	// 		gauge.Observe(
	// 			ctx,
	// 			int64(count),
	// 			attribute.String("pool", pool.Name),
	// 			attribute.String("status", status),
	// 		)
	// 	}
	// }

	// if err = pusher.Stop(ctx); err != nil {
	// 	panic(err)
	// }
}
