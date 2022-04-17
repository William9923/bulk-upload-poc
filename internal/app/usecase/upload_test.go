package usecase

import (
	"testing"

	"github.com/stretchr/testify/mock"
)

type MockFile struct {
	mock.Mock
}

func (mf *MockFile) Read(p []byte) (n int, err error) {
	return 0, nil
}

func TestUsecase_UploadWhitelists(t *testing.T) {

	_ = MockFile{}

	t.Run("should be able to check either file valid / not", func(t *testing.T) {})

	t.Run("should be able to validate data", func(t *testing.T) {
		t.Run("validate failed parsing", func(t *testing.T) {})

		t.Run("validate unknown user", func(t *testing.T) {})

		t.Run("validate invalid status", func(t *testing.T) {})

		t.Run("validate correct instance", func(t *testing.T) {})
	})

	t.Run("should be able to handle failed saving reports file", func(t *testing.T) {})

	t.Run("should be able to handle failed saving reports info to datastore", func(t *testing.T) {})

	t.Run("should be able to return reports information if upload process successful", func(t *testing.T) {})

}
