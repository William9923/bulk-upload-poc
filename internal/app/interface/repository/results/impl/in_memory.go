package resultsrepoimpl

import (
	"fmt"
	"path/filepath"

	"github.com/William9923/bulk-upload-poc/internal/app/domain"
	resultsrepo "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/results"
	"github.com/William9923/bulk-upload-poc/pkg/csv"
)

type InMemoryResultsRepo struct {
	results    []domain.Result
	csvBuilder csv.ICsvBuilder
}

var ErrNoData = fmt.Errorf("no data found")
var ErrTxFailed = fmt.Errorf("failed database transaction")

func NewInMemoryResultsRepo() resultsrepo.IResultsRepo {

	return &InMemoryResultsRepo{
		results:    []domain.Result{},
		csvBuilder: &csv.CsvBuilder{},
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

	// some hardcoded stuff, please change for your own database implementation...
	result.Id = int64(len(impl.results))

	impl.results = append(impl.results, result)
	return result.Id, nil
}

func (impl *InMemoryResultsRepo) SaveResult(result domain.Result) (domain.Result, error) {
	id := int64(len(impl.results))
	url := filepath.Join("data", fmt.Sprintf("result-%d.csv", id))

	// save to file
	header := []string{
		"Id",
		"Name",
		"Status",
		"Upload Result",
		"Reason",
	}

	contents := [][]string{}
	for _, instance := range result.Instances {
		content := make([]string, len(header))
		content[0] = fmt.Sprintf("%d", instance.Data.Id)
		content[1] = string(instance.Data.Name)
		content[2] = fmt.Sprintf("%d", instance.Data.Status)
		content[3] = fmt.Sprintf("%d", instance.Status)
		content[4] = string(instance.Reason)

		contents = append(contents, content)
	}

	csvFile := impl.csvBuilder.Build(header, contents)
	csvFile.Save(url)

	result.URL = url
	return result, nil
}
