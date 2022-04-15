package usecase

import (
	"bytes"
	"context"
	"io"
)

type IUsecase interface {
	ShowUsers(ctx context.Context) (UsersDTO, error)
	ShowResults(ctx context.Context) (ResultsDTO, error)
	ExportUsers(ctx context.Context) (bytes.Buffer, error)
	ExportResult(ctx context.Context, id int64) (bytes.Buffer, error) // TODO: can create additional class for params if the parameter start to be complicated...
	UploadWhitelists(ctx context.Context, file io.Reader) (ResultDTO, error)
}
