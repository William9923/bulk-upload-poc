package usecase

import (
	"path/filepath"
	"testing"

	"github.com/William9923/bulk-upload-poc/internal/app/constant"
	"github.com/William9923/bulk-upload-poc/internal/app/domain"
	mockresultsrp "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/results/mocks"
	mockusersrp "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/users/mocks"
	mockcsv "github.com/William9923/bulk-upload-poc/pkg/csv/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

// Behavior driven unit test
func TestUsecase_ExportUsers(t *testing.T) {

	t.Run("should handle failed to fetch data", func(t *testing.T) {
		// Arrange
		mockUsersRepo := &mockusersrp.IUsersRepo{}
		mockUsersRepo.On("GetUsers").
			Return(nil, errDummy)

		u := NewUsecase(mockUsersRepo, mockResultsRepo, mockCsvBuilder)

		// Act
		got, err := u.ExportUsers(dummyCtx)

		// Assert
		assert.NotNil(t, err, "error should not nil")
		assert.ObjectsAreEqualValues(got, EMPTYFILE)
	})

	t.Run("should handle in case failed to export into file", func(t *testing.T) {
		// Arrange
		mockCsvBuilder := mockcsv.NewBuilder(errDummy)

		u := NewUsecase(mockUsersRepo, mockResultsRepo, mockCsvBuilder)

		// Act
		got, err := u.ExportUsers(dummyCtx)

		// Assert
		assert.NotNil(t, err, "error should not nil")
		assert.ObjectsAreEqualValues(got, EMPTYFILE)

	})

	t.Run("should build correct files", func(t *testing.T) {
		// Arrange
		users := []domain.User{
			{
				Id:     1,
				Name:   "User 1",
				Status: constant.WHITELIST,
			},
			{
				Id:     2,
				Name:   "User 2",
				Status: constant.BLACKLIST,
			},
			{
				Id:     3,
				Name:   "User 3",
				Status: constant.WHITELIST,
			},
		}

		mockUsersRepo := &mockusersrp.IUsersRepo{}
		mockUsersRepo.On("GetUsers").
			Return(users, nil)

		mockCsv := &mockcsv.ICsv{}
		mockCsv.On("Export").
			Return(EMPTYFILE, nil)

		mockCsvBuilder := &mockcsv.ICsvBuilder{}
		mockCsvBuilder.On("Build", mock.Anything, mock.Anything).
			Run(func(args mock.Arguments) {
				// Stub
				arg0 := args.Get(0).([]string)
				arg1 := args.Get(1).([][]string)

				header := []string{
					"Id",
					"Name",
					"Status",
				}

				firstRow := []string{
					"1",
					"User 1",
					"1",
				}

				// Assert
				assert.EqualValues(t, header, arg0)
				assert.EqualValues(t, firstRow, arg1[0])
			}).
			Return(mockCsv, nil)

		u := NewUsecase(mockUsersRepo, mockResultsRepo, mockCsvBuilder)

		// Act
		_, err := u.ExportUsers(dummyCtx)

		// Assert
		assert.Nil(t, err, "error should be nil")
	})

}

func TestUsecase_ExportResult(t *testing.T) {
	t.Run("should handle failed to fetch data", func(t *testing.T) {
		// Arrange
		mockResultsRepo := &mockresultsrp.IResultsRepo{}
		mockResultsRepo.On("GetResult", mock.Anything).
			Return(domain.NullResult(), errDummy)

		u := NewUsecase(mockUsersRepo, mockResultsRepo, mockCsvBuilder)

		// Act
		got, err := u.ExportResult(dummyCtx, 1)

		// Assert
		assert.NotNil(t, err, "error should not nil")
		assert.ObjectsAreEqualValues(got, EMPTYFILE)
	})

	t.Run("should handle in case failed to export into file", func(t *testing.T) {
		// Arrange
		mockCsvBuilder := mockcsv.NewBuilder(errDummy)

		u := NewUsecase(mockUsersRepo, mockResultsRepo, mockCsvBuilder)

		// Act
		got, err := u.ExportResult(dummyCtx, 1)

		// Assert
		assert.NotNil(t, err, "error should not nil")
		assert.ObjectsAreEqualValues(got, EMPTYFILE)

	})

	t.Run("should build correct files", func(t *testing.T) {
		// Arrange
		userData := domain.NullUser()
		result := domain.Result{
			Id: 1,
			Instances: []domain.UploadInstance{
				{
					Idx:    0,
					Data:   userData,
					Status: constant.SUCCESS,
					Reason: "-",
				},
				{
					Idx:    1,
					Data:   userData,
					Status: constant.FAILED,
					Reason: constant.FAILEDPARSING,
				},
				{
					Idx:    2,
					Data:   userData,
					Status: constant.NOTPROCESSED,
				},
			},
			URL: filepath.Join("data", "tmp.csv"),
		}

		mockResultsRepo := &mockresultsrp.IResultsRepo{}
		mockResultsRepo.On("GetResult", mock.Anything).
			Return(result, nil)

		mockCsv := &mockcsv.ICsv{}
		mockCsv.On("Export").
			Return(EMPTYFILE, nil)

		mockCsvBuilder := &mockcsv.ICsvBuilder{}
		mockCsvBuilder.On("Build", mock.Anything, mock.Anything).
			Run(func(args mock.Arguments) {
				// Stub
				arg0 := args.Get(0).([]string)
				arg1 := args.Get(1).([][]string)

				header := []string{
					"Id",
					"Name",
					"Status",
					"Upload Result",
					"Reason",
				}

				firstRow := []string{
					"-1",
					"null-user",
					"0",
					"1",
					"-",
				}

				secondRow := []string{
					"-1",
					"null-user",
					"0",
					"0",
					constant.FAILEDPARSING,
				}

				// Assert
				assert.EqualValues(t, header, arg0)
				assert.EqualValues(t, firstRow, arg1[0])
				assert.EqualValues(t, secondRow, arg1[1])
			}).
			Return(mockCsv, nil)

		u := NewUsecase(mockUsersRepo, mockResultsRepo, mockCsvBuilder)

		// Act
		_, err := u.ExportResult(dummyCtx, 1)

		// Assert
		assert.Nil(t, err, "error should be nil")
	})
}
