package resultsrepo

import (
	"github.com/William9923/bulk-upload-poc/internal/app/domain"
)

type IResultsRepo interface {
	GetResult(id int64) (domain.Result, error)
	GetResults() ([]domain.Result, error)
	CreateResult(domain.Result) (int64, error)
	SaveResult(domain.Result) (domain.Result, error)
}
