package main

import "github.com/gozelle/log"

func startOven(degree int) {
	log.Helper()
	log.Info("Starting oven", "degree", degree)
}

func main() {
	log.SetReportCaller(true)
	startOven(400)
}
