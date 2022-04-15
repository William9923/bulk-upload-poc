package usecase

import "github.com/William9923/bulk-upload-poc/internal/app/domain"

type UsersDTO struct {
	Users []domain.User `json:"users"`
}

type ResultsDTO struct {
	Results []domain.Result `json:"results"`
}

type ResultDTO struct {
	Id  int64  `json:"id"`
	URL string `json:"url"`
}
