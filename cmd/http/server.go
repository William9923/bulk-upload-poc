package main

import (
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"gopkg.in/tylerb/graceful.v1"
)

// Notes: configure in seperate configuration files for bigger projects...
const (
	HTTPPORT = "localhost:8080"
	TIMEOUT  = 10
)

func startServer(r *httprouter.Router) error {

	srv := &graceful.Server{
		Timeout: TIMEOUT * time.Second,

		Server: &http.Server{
			Addr:    HTTPPORT,
			Handler: r,
		},
	}

	srv.ListenAndServe()
	return nil
}
