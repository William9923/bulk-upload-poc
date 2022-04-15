package http

import (
	"net/http"
	"strconv"

	"github.com/julienschmidt/httprouter"

	"github.com/William9923/bulk-upload-poc/internal/app/usecase"
	"github.com/William9923/bulk-upload-poc/pkg/response"
	"github.com/sirupsen/logrus"
)

type HTTPHandler interface {
	HandleShowUsers(w http.ResponseWriter, req *http.Request, p httprouter.Params)
	HandleShowUploadResults(w http.ResponseWriter, req *http.Request, p httprouter.Params)
	HandleBulkUploadWhitelists(w http.ResponseWriter, req *http.Request, p httprouter.Params)
	HandleExportUsers(w http.ResponseWriter, req *http.Request, p httprouter.Params)
	HandleExportUploadResult(w http.ResponseWriter, req *http.Request, p httprouter.Params)
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
	ctx := req.Context()
	logger := logrus.WithContext(ctx)
	resp := response.NewJSONResponse(ctx)
	defer resp.Render(w)

	logger.Info("begin showing users...")

	res, err := h.usecase.ShowUsers(ctx)
	if err != nil {
		logger.Info("error when showing users : ", err.Error())
		resp.SetError(http.StatusInternalServerError, err)
		return
	}

	resp.SetData(res)

}

func (h handler) HandleShowUploadResults(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	ctx := req.Context()
	logger := logrus.WithContext(ctx)
	resp := response.NewJSONResponse(ctx)
	defer resp.Render(w)

	logger.Info("begin showing users...")

	res, err := h.usecase.ShowResults(ctx)
	if err != nil {
		logger.Info("error when showing upload results : ", err.Error())
		resp.SetError(http.StatusInternalServerError, err)
		return
	}

	resp.SetData(res)
}

func (h handler) HandleBulkUploadWhitelists(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	ctx := req.Context()
	logger := logrus.WithContext(ctx)
	resp := response.NewJSONResponse(ctx)
	defer resp.Render(w)

	logger.Info("begin uploading whitelists...")

	file, _, err := req.FormFile("file")
	if err != nil {
		logger.Info("error retrieve file from form file : ", err.Error())
		resp.SetError(http.StatusBadRequest, err)
		return
	}
	res, err := h.usecase.UploadWhitelists(ctx, file)
	if err != nil {
		logger.Info("error when uploading whitelists : ", err.Error())
		resp.SetError(http.StatusInternalServerError, err)
		return
	}

	resp.SetData(res)

}

func (h handler) HandleExportUsers(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	ctx := req.Context()
	logger := logrus.WithContext(ctx)
	resp := response.NewFileResponse(ctx)
	defer resp.Render(w)

	logger.Info("begin exporting users...")

	buffer, err := h.usecase.ExportUsers(ctx)
	if err != nil {
		logger.Info("error when exporting users : ", err.Error())
		resp.SetError(http.StatusInternalServerError, err)
		return
	}

	resp.SetData(buffer)
}

func (h handler) HandleExportUploadResult(w http.ResponseWriter, req *http.Request, p httprouter.Params) {
	ctx := req.Context()
	logger := logrus.WithContext(ctx)
	resp := response.NewFileResponse(ctx)
	defer resp.Render(w)

	logger.Info("begin exporting upload results...")

	id, err := strconv.Atoi(p.ByName("id"))
	if err != nil {
		logger.Info("failed to parse input : ", err.Error())
		resp.SetError(http.StatusBadRequest, err)
	}

	buffer, err := h.usecase.ExportResult(ctx, int64(id))
	if err != nil {
		logger.Info("error when exporting upload results : ", err.Error())
		resp.SetError(http.StatusInternalServerError, err)
		return
	}

	resp.SetData(buffer)
}
