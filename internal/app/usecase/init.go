package usecase

import (
	resultsrepo "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/results"
	usersrepo "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/users"
)

type Usecase struct {
	usersRepo   usersrepo.IUsersRepo
	resultsRepo resultsrepo.IResultsRepo
}

func NewUsecase(
	usersRepo usersrepo.IUsersRepo,
	resultsRepo resultsrepo.IResultsRepo,
) IUsecase {
	return &Usecase{
		usersRepo:   usersRepo,
		resultsRepo: resultsRepo,
	}
}

func (u Usecase) ShowUsers() (UsersDTO, error) {
	data, err := u.usersRepo.GetUsers()
	if err != nil {
		return UsersDTO{}, err
	}

	dtos := make([]UserDTO, len(data))
	for i, datum := range data {
		dtos[i] = UserDTO{
			User: datum,
		}
	}
	return UsersDTO{users: dtos}, nil

}
func (u Usecase) ShowResults() (ResultsDTO, error) {
	data, err := u.resultsRepo.GetResults()
	if err != nil {
		return ResultsDTO{}, err
	}

	dtos := make([]ResultDTO, len(data))
	for i, datum := range data {
		dtos[i] = ResultDTO{
			Result: datum,
		}
	}
	return ResultsDTO{results: dtos}, nil
}
