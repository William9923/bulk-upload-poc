package main

import (
	"flag"
	"os"

	log "github.com/sirupsen/logrus"
)

func main() {

	// init logger
	// get flag when running the service
	env := flag.String("env", "dev", "environment for running the service")
	if *env == "prod" {
		log.SetFormatter(&log.JSONFormatter{})
	}
	log.SetOutput(os.Stdout)

	// initialize dependencies & start application
	log.Info("starting service...")
	startApp()
}
