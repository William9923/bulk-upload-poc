package resultsrepoimpl

import (
	"github.com/William9923/bulk-upload-poc/internal/app/domain"
	resultsrepo "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/results"
)

type InMemoryResultsRepo struct {
	results []domain.Result
}

func NewInMemoryResultsRepo() resultsrepo.IResultsRepo {
	return &InMemoryResultsRepo{
		results: []domain.Result{},
	}
}

func (impl InMemoryResultsRepo) GetResult(id int64) domain.Result {
	for _, result := range impl.results {
		if result.Id == id {
			return result
		}
	}
	return domain.NullResult()
}

func (impl InMemoryResultsRepo) GetResults() []domain.Result {
	return impl.results
}
