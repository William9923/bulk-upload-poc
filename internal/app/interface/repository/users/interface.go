package usersrepo

import (
	"github.com/William9923/bulk-upload-poc/internal/app/domain"
)

type IUsersRepo interface {
	GetUser(id int64) domain.User
	GetUsers() []domain.User
}
