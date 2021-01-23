package main

import (
	"flag"
	"os"
)

type MonitoringEnv struct {
	MonitoringBaseUrl string
}

func LoadEnv() (env MonitoringEnv) {
	flag.StringVar(&(env.MonitoringBaseUrl), "monitoringBaseUrl", os.Getenv("MONITORING_BASE_URL"), "the monitoring Api base URL")
	flag.Parse()

	return
}
