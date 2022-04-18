package usersrepoimpl

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/William9923/bulk-upload-poc/internal/app/constant"
	"github.com/William9923/bulk-upload-poc/internal/app/domain"
	usersrepo "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/users"
)

var ErrNoData = fmt.Errorf("no data found")
var ErrTxFailed = fmt.Errorf("failed database transaction")

const NOTFOUND = int64(-1)

type InMemoryUsersRepo struct {
	users []domain.User
}

func NewInMemoryResultsRepo() usersrepo.IUsersRepo {

	return &InMemoryUsersRepo{
		users: seeding(),
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

func (impl *InMemoryUsersRepo) UpdateUsers(users []domain.User) []error {

	// 1. Validate all bulk change data ...
	indexes := []int64{}
	errors := make([]error, len(users))
	for i, user := range users {
		idx := impl.findIdx(user, impl.users)
		if idx == NOTFOUND {
			errors[i] = ErrNoData
		}
		indexes = append(indexes, idx)
	}

	// 2. Update all....
	for i, idx := range indexes {

		if err := randomError(); errors[i] == nil && err != nil {
			errors[i] = ErrTxFailed
		}

		if errors[i] == nil {
			impl.users[idx] = users[i]
		}

	}

	return errors
}

func (impl InMemoryUsersRepo) findIdx(user domain.User, users []domain.User) int64 {
	for idx, u := range users {
		if u.Id == user.Id {
			return int64(idx)
		}
	}
	return NOTFOUND
}

func seeding() []domain.User {

	users := make([]domain.User, 100000)

	for i := range users {

		status := constant.WHITELIST
		if status%2 == 0 {
			status = constant.BLACKLIST
		}

		users[i] = domain.User{
			Id:     int64(i + 1),
			Name:   fmt.Sprintf("User %d", i+1),
			Status: int64(status),
		}
	}

	return users
}

func randomError() error {
	rand.Seed(time.Now().UnixNano())
	if rand.Intn(2) == 1 {
		return ErrTxFailed
	}
	return nil
}
