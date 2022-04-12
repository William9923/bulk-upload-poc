package main

import (
	"flag"
	"os"

	"github.com/sirupsen/logrus"
)

func main() {

	log := logrus.New()
	// init logger
	env := flag.String("env", "dev", "environment for running the service")
	log.SetFormatter(&logrus.TextFormatter{
		DisableColors:   false,
		TimestampFormat: "2006-01-02 15:04:05",
		FullTimestamp:   true,
	})
	if *env == "prod" {
		log.SetFormatter(&logrus.JSONFormatter{
			TimestampFormat: "2006-01-02 15:04:05",
		})
	}
	log.SetOutput(os.Stdout)

	// initialize dependencies & start application
	log.Info("starting service...")
	startApp()
}
