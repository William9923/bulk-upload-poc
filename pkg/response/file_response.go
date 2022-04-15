package response

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/sirupsen/logrus"
)

var ErrInvalidConversion = fmt.Errorf("not file buffer")

// FileResponse defines file output for http response for the client
type FileResponse struct {
	Data       bytes.Buffer    `json:"-"`
	Error      string          `json:"error,omitempty"`
	Context    context.Context `json:"-"`
	StatusCode int             `json:"-"`
}

func NewFileResponse(ctx context.Context) IResponse {
	logger := logrus.WithContext(ctx)

	logger.Infof("start new http request...")
	return &FileResponse{
		Context: ctx,
	}
}

func (r *FileResponse) SetData(data interface{}) {

	switch data.(type) {
	case bytes.Buffer:
		r.Data = data.(bytes.Buffer)
	default:
		r.Error = ErrInvalidConversion.Error()
	}
}

func (r *FileResponse) SetError(code int, err error) {
	r.StatusCode = code
	r.Error = err.Error()
}

func (r *FileResponse) Render(w http.ResponseWriter) {
	logger := logrus.WithContext(r.Context)
	statusCode := http.StatusOK
	if r.StatusCode != 0 {
		statusCode = r.StatusCode
	}

	w.Header().Set("Content-Disposition", "attachment; filename=user_block_unblock.xlsx")
	w.Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	w.WriteHeader(statusCode)
	_, err := r.Data.WriteTo(w)
	if err != nil {
		r.Error = err.Error()
	}

	// Handle if error exist / happen
	if r.Error != "" {
		contentType := "application/json"
		response, err := json.Marshal(r)
		if err != nil {
			logger.Error("failed to marshal response. err: ", err)
			r.StatusCode = http.StatusInternalServerError
		}
		logger.Infof("request completed with header status code %d", statusCode)

		// write response
		w.Header().Set("Content-Type", contentType)
		w.Write(response)
	}

	return

}
