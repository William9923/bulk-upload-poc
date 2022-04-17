package usecase

import (
	"context"

	resultsrepo "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/results"
	usersrepo "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/users"
	"github.com/William9923/bulk-upload-poc/pkg/csv"
)

type Usecase struct {
	usersRepo   usersrepo.IUsersRepo
	resultsRepo resultsrepo.IResultsRepo
	csvBuilder  csv.ICsvBuilder
}

func NewUsecase(
	usersRepo usersrepo.IUsersRepo,
	resultsRepo resultsrepo.IResultsRepo,
	csvBuilder csv.ICsvBuilder,
) IUsecase {
	return &Usecase{
		usersRepo:   usersRepo,
		resultsRepo: resultsRepo,
		csvBuilder:  csvBuilder,
	}
}

func (u Usecase) ShowUsers(ctx context.Context) (UsersDTO, error) {
	users, err := u.usersRepo.GetUsers()
	if err != nil {
		return UsersDTO{}, err
	}

	return UsersDTO{Users: users}, nil

}

func (u Usecase) ShowResults(ctx context.Context) (ResultsDTO, error) {
	results, err := u.resultsRepo.GetResults()
	if err != nil {
		return ResultsDTO{}, err
	}

	return ResultsDTO{Results: results}, nil
}
