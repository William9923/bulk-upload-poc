package usecase

type IUsecase interface {
	ShowUsers() (UsersDTO, error)
	ShowResults() (ResultsDTO, error)
}
