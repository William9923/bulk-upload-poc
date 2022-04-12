package http

import (
	"net/http"

	"github.com/julienschmidt/httprouter"

	"github.com/William9923/bulk-upload-poc/internal/app/usecase"
)

type HTTPHandler interface {
	HandleShowUsers(w http.ResponseWriter, req *http.Request, p httprouter.Params)
	HandleShowUploadResults(w http.ResponseWriter, req *http.Request, p httprouter.Params)
}

// Notes: Please seperate usecase if you want to expand the capability.
// How you split the usecase is based on your own convention, feel free to decide by yourself on how to split the usecase
type handler struct {
	usecase usecase.IUsecase
}

// New test constructor
func NewHTTPHandler(uc usecase.IUsecase) HTTPHandler {
	return &handler{
		usecase: uc,
	}
}

// Since the deliveries is not complicated, we will write in 1 file.
// Notes, if the logic in deliveries is too long, please split into multiple files or use simpler logic (move to the usecase parts)
func (h handler) HandleShowUsers(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	panic("implement this shid!")
}

func (h handler) HandleShowUploadResults(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	panic("implement this shid!")
}
