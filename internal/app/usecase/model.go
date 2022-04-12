package usecase

import "github.com/William9923/bulk-upload-poc/internal/app/domain"

type UserDTO struct {
	domain.User
}

type UsersDTO struct {
	users []UserDTO
}

type ResultDTO struct {
	domain.Result
}
type ResultsDTO struct {
	results []ResultDTO
}
