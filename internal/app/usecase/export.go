package usecase

import "context"

func (u Usecase) ExportUsers(ctx context.Context) ([]byte, error) {
	return make([]byte, 0), nil
}
func (u Usecase) ExportResult(ctx context.Context, id int64) ([]byte, error) {
	return make([]byte, 0), nil
}
