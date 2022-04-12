package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"
	"github.com/sirupsen/logrus"
)

func Logging(h httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {
		ctx := r.Context()
		logger := logrus.WithContext(ctx)
		logger.Println("start http request...")
		h(w, r, ps)
	}
}

func Authentication(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		ctx := req.Context()
		logger := logrus.WithContext(ctx)
		logger.Println("start authentication...")
		fn(w, req, p)
	}
}

func Authorization(fn httprouter.Handle) httprouter.Handle {
	return func(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
		ctx := req.Context()
		logger := logrus.WithContext(ctx)
		logger.Println("checking permission...")
		fn(w, req, p)
	}
}
