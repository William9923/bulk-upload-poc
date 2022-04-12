package main

import (
	"net/http"

	handler "github.com/William9923/bulk-upload-poc/internal/app/interface/handler/http"
	"github.com/William9923/httpmiddleware"
	"github.com/julienschmidt/httprouter"
)

func Pong(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("pong"))
}

func routes() *httprouter.Router {

	router := httprouter.New()
	middleware := httpmiddleware.New()
	middleware.Use(handler.Logging)

	router.GET("/ping", middleware.Wrap(Pong))

	return router
}
