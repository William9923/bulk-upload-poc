package mocks

import (
	"github.com/William9923/bulk-upload-poc/internal/app/domain"
	"github.com/stretchr/testify/mock"
)

func New(err error) *IUsersRepo {

	mockUsersRepo := &IUsersRepo{}
	mockUsersRepo.On("GetUser", mock.Anything).Return(domain.NullUser(), err)
	mockUsersRepo.On("GetUsers").Return(make([]domain.User, 0), err)

	errs := []error{}
	if err != nil {
		errs = append(errs, err)
	}

	mockUsersRepo.On("UpdateUsers", mock.Anything).Return(errs)
	return mockUsersRepo
}
