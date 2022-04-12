package usersrepoimpl

import (
	"fmt"

	"github.com/William9923/bulk-upload-poc/internal/app/domain"
	usersrepo "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/users"
)

type noData error

var ErrNoData = noData(fmt.Errorf("no data found"))

type InMemoryUsersRepo struct {
	users []domain.User
}

func NewInMemoryResultsRepo() usersrepo.IUsersRepo {
	return &InMemoryUsersRepo{
		users: []domain.User{},
	}
}

func (impl InMemoryUsersRepo) GetUser(id int64) (domain.User, error) {
	for _, result := range impl.users {
		if result.Id == id {
			return result, nil
		}
	}
	return domain.NullUser(), ErrNoData
}

func (impl InMemoryUsersRepo) GetUsers() ([]domain.User, error) {
	return impl.users, nil
}
