package response

import (
	"context"
	"encoding/json"
	"net/http"

	"github.com/sirupsen/logrus"
)

type IResponse interface {
	Render(w http.ResponseWriter)
	SetError(code int, err error)
}

// Response defines http response for the client
type JSONResponse struct {
	Data       interface{}     `json:"data,omitempty"`
	Error      string          `json:"error,omitempty"`
	Context    context.Context `json:"-"`
	StatusCode int             `json:"-"`
}

func NewJSONResponse(ctx context.Context) *JSONResponse {
	logger := logrus.WithContext(ctx)

	logger.Infof("start new http request...")
	return &JSONResponse{
		Context: ctx,
	}
}

func (r *JSONResponse) SetError(code int, err error) {
	r.StatusCode = code
	r.Error = err.Error()
}

func (r *JSONResponse) Render(w http.ResponseWriter) {
	logger := logrus.WithContext(r.Context)
	statusCode := http.StatusOK
	if r.StatusCode != 0 {
		statusCode = r.StatusCode
	}
	contentType := "application/json"

	if r.Error != "" {
		r.Data = nil
	}

	response, err := json.Marshal(r)
	if err != nil {
		logger.Error("failed to marshal response. err: ", err)
		r.StatusCode = http.StatusInternalServerError
	}
	logger.Infof("request completed with header status code %d", statusCode)

	// write response
	w.Header().Set("Content-Type", contentType)
	w.WriteHeader(statusCode)
	w.Write(response)
}
