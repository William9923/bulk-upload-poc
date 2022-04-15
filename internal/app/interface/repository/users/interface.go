package usersrepo

import (
	"github.com/William9923/bulk-upload-poc/internal/app/domain"
)

type IUsersRepo interface {
	GetUser(id int64) (domain.User, error)
	GetUsers() ([]domain.User, error)
	UpdateUsers(users []domain.User) []error
}
