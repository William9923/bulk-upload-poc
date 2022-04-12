package domain

import (
	"github.com/William9923/bulk-upload-poc/internal/app/constant"
)

type User struct {
	Id     int64
	Name   string
	Status int64
}

func NullUser() User {
	return User{
		Id:     -1,
		Name:   "null-user",
		Status: constant.BLACKLIST,
	}
}
