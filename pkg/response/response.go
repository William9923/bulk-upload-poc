package response

import (
	"net/http"
)

type IResponse interface {
	Render(w http.ResponseWriter)
	SetError(code int, err error)
	SetData(data interface{})
}
