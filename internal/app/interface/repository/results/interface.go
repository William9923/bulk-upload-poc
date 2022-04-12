package resultsrepo

import (
	"github.com/William9923/bulk-upload-poc/internal/app/domain"
)

type IResultsRepo interface {
	GetResult(id int64) domain.Result
	GetResults() []domain.Result
}
