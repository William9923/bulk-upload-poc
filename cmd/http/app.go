package main

import (
	httpHandler "github.com/William9923/bulk-upload-poc/internal/app/interface/handler/http"
	resultsrepoimpl "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/results/impl"
	usersrepoimpl "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/users/impl"
	uc "github.com/William9923/bulk-upload-poc/internal/app/usecase"
	"github.com/William9923/bulk-upload-poc/pkg/csv"
)

func startApp() error {
	// Build infrastructure
	usersRepo := usersrepoimpl.NewInMemoryResultsRepo()
	resultsRepo := resultsrepoimpl.NewInMemoryResultsRepo()
	csvBuilder := csv.NewCsvBuilder()

	usecase := uc.NewUsecase(usersRepo, resultsRepo, csvBuilder)

	// init handler
	handler := httpHandler.NewHTTPHandler(usecase)

	// web server & handlers
	r := routes(handler)

	// start server
	return startServer(r)
}
