package resultsrepoimpl

import (
	"fmt"

	"github.com/William9923/bulk-upload-poc/internal/app/domain"
	resultsrepo "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/results"
)

type InMemoryResultsRepo struct {
	results []domain.Result
}

type noData error
var ErrNoData = noData(fmt.Errorf("no data found"))

func NewInMemoryResultsRepo() resultsrepo.IResultsRepo {
	return &InMemoryResultsRepo{
		results: []domain.Result{},
	}
}

func (impl InMemoryResultsRepo) GetResult(id int64) (domain.Result, error) {
	for _, result := range impl.results {
		if result.Id == id {
			return result, nil
		}
	}
	return domain.NullResult(), ErrNoData
}

func (impl InMemoryResultsRepo) GetResults() ([]domain.Result, error) {
	return impl.results, nil
}

func (impl *InMemoryResultsRepo) CreateResult(result domain.Result) (int64, error) {
	result.Id = int64(len(impl.results))
	impl.results = append(impl.results, result)
	return result.Id, nil
}
