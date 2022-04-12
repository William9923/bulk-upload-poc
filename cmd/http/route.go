package main

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func routes() *httprouter.Router {
	r := httprouter.New()

	r.GET("/ping", func(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
		w.Write([]byte("pong"))
	})

	return r
}
