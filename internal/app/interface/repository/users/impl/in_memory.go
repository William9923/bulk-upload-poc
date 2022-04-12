package usersrepoimpl

import (
	"github.com/William9923/bulk-upload-poc/internal/app/domain"
	usersrepo "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/users"
)

type InMemoryUsersRepo struct {
	users []domain.User
}

func NewInMemoryResultsRepo() usersrepo.IUsersRepo {
	return &InMemoryUsersRepo{
		users: []domain.User{},
	}
}

func (impl InMemoryUsersRepo) GetUser(id int64) domain.User {
	for _, result := range impl.users {
		if result.Id == id {
			return result
		}
	}
	return domain.NullUser()
}

func (impl InMemoryUsersRepo) GetUsers() []domain.User {
	return impl.users
}
