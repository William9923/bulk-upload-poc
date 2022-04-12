package http

import (
	"log"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

func Logging(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		logger := log.Default()
		logger.Println("start http request...")
		h(w, r, ps)
	}
}

func Authentication(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		logger := log.Default()
		logger.Println("start authentication...")
		fn(w, req, p)
	}
}

func Authorization(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		logger := log.Default()
		logger.Println("checking permission...")
		fn(w, req, p)
	}
}
