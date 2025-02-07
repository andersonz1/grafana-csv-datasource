package main

import (
	"os"

	"github.com/andersonz1/grafana-plugin-sdk-go/backend/datasource"
	"github.com/andersonz1/grafana-plugin-sdk-go/backend/log"
)

func main() {
	logger := log.New()

	ds := newDataSource(logger)

	opts := datasource.ServeOpts{
		QueryDataHandler:   ds,
		CheckHealthHandler: ds,
	}

	if err := datasource.Serve(opts); err != nil {
		logger.Error(err.Error())
		os.Exit(1)
	}
}
