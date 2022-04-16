package mocks

import (
	domain "github.com/William9923/bulk-upload-poc/internal/app/domain"
	"github.com/stretchr/testify/mock"
)

func New(err error) *IResultsRepo {

	mockUsersRepo := &IResultsRepo{}
	mockUsersRepo.On("CreateResult", mock.Anything).Return(err)
	mockUsersRepo.On("GetResult", mock.Anything).Return(domain.NullResult(), err)
	mockUsersRepo.On("GetResults", mock.Anything).Return(make([]domain.Result, 0), err)
	mockUsersRepo.On("GetResults", mock.Anything).Return(domain.NullResult(), err)

	return mockUsersRepo
}
