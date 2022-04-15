package main

import (
	"net/http"

	httpHandler "github.com/William9923/bulk-upload-poc/internal/app/interface/handler/http"
	"github.com/William9923/httpmiddleware"
	"github.com/julienschmidt/httprouter"
)

func Pong(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	w.Write([]byte("pong"))
}

func routes(handler httpHandler.HTTPHandler) *httprouter.Router {

	router := httprouter.New()
	middleware := httpmiddleware.New()
	middleware.Use(httpHandler.Logging)

	router.GET("/ping", middleware.Wrap(Pong))

	router.GET("/users", middleware.Wrap(handler.HandleShowUsers))
	router.GET("/results", middleware.Wrap(handler.HandleShowUploadResults))
	router.POST("/upload", middleware.Wrap(handler.HandleBulkUploadWhitelists))

	return router
}
