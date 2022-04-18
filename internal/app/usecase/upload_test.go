package usecase

import (
	"reflect"
	"testing"

	"github.com/William9923/bulk-upload-poc/internal/app/constant"
	"github.com/William9923/bulk-upload-poc/internal/app/domain"
	mockresultsrp "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/results/mocks"
	mockusersrp "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/users/mocks"
	mockcsv "github.com/William9923/bulk-upload-poc/pkg/csv/mocks"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type MockFile struct {
	mock.Mock
}

func (mf *MockFile) Read(p []byte) (n int, err error) {
	return 0, nil
}

func TestUsecase_UploadWhitelists(t *testing.T) {

	file := &MockFile{}

	t.Run("should be able to check either file valid / not", func(t *testing.T) {

		// Arrange
		mockCsvBuilder := mockcsv.NewBuilder(errDummy)
		u := NewUsecase(mockUsersRepo, mockResultsRepo, mockCsvBuilder)

		// Act
		got, err := u.UploadWhitelists(dummyCtx, file)

		// Assert
		assert.NotNil(t, err, "error should not nil")
		assert.ObjectsAreEqualValues(got, EMPTYRESULTDTO)
	})

	t.Run("should be able to validate data", func(t *testing.T) {
		t.Run("validate failed parsing", func(t *testing.T) {

			// Arrange
			contents := [][]string{
				{"invalid-number", "User 1", "1"},
				{"2", "User 2", "0"},
			}
			mockCsv := &mockcsv.ICsv{}
			mockCsv.On("GetContents").
				Return(contents, nil)

			mockCsvBuilder := &mockcsv.ICsvBuilder{}
			mockCsvBuilder.On("FromFile", mock.Anything).
				Return(mockCsv, nil)

			mockUsersRepo := &mockusersrp.IUsersRepo{}
			mockUsersRepo.On("GetUser", int64(2)).Return(domain.User{Id: 2, Name: "User 2", Status: constant.BLACKLIST}, nil)
			mockUsersRepo.On("UpdateUsers", mock.Anything).Return([]error{errDummy, nil})

			mockResultsRepo := &mockresultsrp.IResultsRepo{}
			mockResultsRepo.On("SaveResult", mock.Anything).Run(func(args mock.Arguments) {
				// Stub
				arg0 := args.Get(0).(domain.Result)
				uploadInstances := arg0.Instances

				firstInstance := domain.UploadInstance{
					Idx:    0,
					Data:   domain.NullUser(),
					Status: constant.FAILED,
					Reason: constant.FAILEDPARSING,
				}

				secondInstance := domain.UploadInstance{
					Idx: 1,
					Data: domain.User{
						Id:     2,
						Name:   "User 2",
						Status: constant.BLACKLIST,
					},
					Status: constant.SUCCESS,
					Reason: constant.VALID,
				}

				// Assert
				if !reflect.DeepEqual(uploadInstances[0], firstInstance) {
					t.Errorf("False first upload instance when failed parsing occurs! got = %v, want %v", uploadInstances[0], firstInstance)
				}
				if !reflect.DeepEqual(uploadInstances[1], secondInstance) {
					t.Errorf("False second upload instance when failed parsing occurs! got = %v, want %v", uploadInstances[1], secondInstance)
				}
			}).Return(domain.NullResult(), nil)

			mockResultsRepo.On("CreateResult", mock.Anything).Return(int64(1), nil)

			u := NewUsecase(mockUsersRepo, mockResultsRepo, mockCsvBuilder)

			// Act
			_, err := u.UploadWhitelists(dummyCtx, file)

			// Assert
			assert.Nil(t, err)
		})

		t.Run("validate unknown user", func(t *testing.T) {
			// Arrange
			contents := [][]string{
				{"1", "User 1", "1"},
				{"2", "User 2", "0"},
			}
			mockCsv := &mockcsv.ICsv{}
			mockCsv.On("GetContents").
				Return(contents, nil)

			mockCsvBuilder := &mockcsv.ICsvBuilder{}
			mockCsvBuilder.On("FromFile", mock.Anything).
				Return(mockCsv, nil)

			mockUsersRepo := &mockusersrp.IUsersRepo{}
			mockUsersRepo.On("GetUser", int64(2)).Return(domain.User{Id: 2, Name: "User 2", Status: constant.BLACKLIST}, nil)
			mockUsersRepo.On("GetUser", int64(1)).Return(domain.NullUser(), errDummy)
			mockUsersRepo.On("UpdateUsers", mock.Anything).Return([]error{errDummy, nil})

			mockResultsRepo := &mockresultsrp.IResultsRepo{}
			mockResultsRepo.On("SaveResult", mock.Anything).Run(func(args mock.Arguments) {
				// Stub
				arg0 := args.Get(0).(domain.Result)
				uploadInstances := arg0.Instances

				firstInstance := domain.UploadInstance{
					Idx: 0,
					Data: domain.User{
						Id:     1,
						Name:   "User 1",
						Status: constant.BLACKLIST,
					},
					Status: constant.FAILED,
					Reason: constant.USERNOTFOUND,
				}

				secondInstance := domain.UploadInstance{
					Idx: 1,
					Data: domain.User{
						Id:     2,
						Name:   "User 2",
						Status: constant.BLACKLIST,
					},
					Status: constant.SUCCESS,
					Reason: constant.VALID,
				}

				// Assert
				if !reflect.DeepEqual(uploadInstances[0], firstInstance) {
					t.Errorf("False first upload instance when failed parsing occurs! got = %v, want %v", uploadInstances[0], firstInstance)
				}
				if !reflect.DeepEqual(uploadInstances[1], secondInstance) {
					t.Errorf("False second upload instance when failed parsing occurs! got = %v, want %v", uploadInstances[1], secondInstance)
				}
			}).Return(domain.NullResult(), nil)

			mockResultsRepo.On("CreateResult", mock.Anything).Return(int64(1), nil)

			u := NewUsecase(mockUsersRepo, mockResultsRepo, mockCsvBuilder)

			// Act
			_, err := u.UploadWhitelists(dummyCtx, file)

			// Assert
			assert.Nil(t, err)
		})

		t.Run("validate invalid status", func(t *testing.T) {
			// Arrange
			contents := [][]string{
				{"1", "User 1", "2"},
				{"2", "User 2", "0"},
			}
			mockCsv := &mockcsv.ICsv{}
			mockCsv.On("GetContents").
				Return(contents, nil)

			mockCsvBuilder := &mockcsv.ICsvBuilder{}
			mockCsvBuilder.On("FromFile", mock.Anything).
				Return(mockCsv, nil)

			mockUsersRepo := &mockusersrp.IUsersRepo{}
			mockUsersRepo.On("GetUser", int64(1)).Return(domain.User{Id: 1, Name: "User 1", Status: constant.BLACKLIST}, nil)
			mockUsersRepo.On("GetUser", int64(2)).Return(domain.User{Id: 2, Name: "User 2", Status: constant.WHITELIST}, nil)
			mockUsersRepo.On("UpdateUsers", mock.Anything).Return([]error{errDummy, nil})

			mockResultsRepo := &mockresultsrp.IResultsRepo{}
			mockResultsRepo.On("SaveResult", mock.Anything).Run(func(args mock.Arguments) {
				// Stub
				arg0 := args.Get(0).(domain.Result)
				uploadInstances := arg0.Instances

				firstInstance := domain.UploadInstance{
					Idx: 0,
					Data: domain.User{
						Id:     1,
						Name:   "User 1",
						Status: constant.BLACKLIST,
					},
					Status: constant.FAILED,
					Reason: constant.STATUSNOTAVAILABLE,
				}

				secondInstance := domain.UploadInstance{
					Idx: 1,
					Data: domain.User{
						Id:     2,
						Name:   "User 2",
						Status: constant.BLACKLIST,
					},
					Status: constant.SUCCESS,
					Reason: constant.VALID,
				}

				// Assert
				if !reflect.DeepEqual(uploadInstances[0], firstInstance) {
					t.Errorf("False first upload instance when failed parsing occurs! got = %v, want %v", uploadInstances[0], firstInstance)
				}
				if !reflect.DeepEqual(uploadInstances[1], secondInstance) {
					t.Errorf("False second upload instance when failed parsing occurs! got = %v, want %v", uploadInstances[1], secondInstance)
				}
			}).Return(domain.NullResult(), nil)

			mockResultsRepo.On("CreateResult", mock.Anything).Return(int64(1), nil)

			u := NewUsecase(mockUsersRepo, mockResultsRepo, mockCsvBuilder)

			// Act
			_, err := u.UploadWhitelists(dummyCtx, file)

			// Assert
			assert.Nil(t, err)
		})

	})

	t.Run("should be able to handle failed when updating users", func(t *testing.T) {
		// Arrange
		contents := [][]string{
			{"1", "User 1", "2"},
			{"2", "User 2", "0"},
		}
		mockCsv := &mockcsv.ICsv{}
		mockCsv.On("GetContents").
			Return(contents, nil)

		mockCsvBuilder := &mockcsv.ICsvBuilder{}
		mockCsvBuilder.On("FromFile", mock.Anything).
			Return(mockCsv, nil)

		mockUsersRepo := &mockusersrp.IUsersRepo{}
		mockUsersRepo.On("GetUser", int64(1)).Return(domain.User{Id: 1, Name: "User 1", Status: constant.BLACKLIST}, nil)
		mockUsersRepo.On("GetUser", int64(2)).Return(domain.User{Id: 2, Name: "User 2", Status: constant.WHITELIST}, nil)
		mockUsersRepo.On("UpdateUsers", mock.Anything).Return([]error{errDummy, errDummy})

		mockResultsRepo := &mockresultsrp.IResultsRepo{}
		mockResultsRepo.On("SaveResult", mock.Anything).Run(func(args mock.Arguments) {
			// Stub
			arg0 := args.Get(0).(domain.Result)
			uploadInstances := arg0.Instances

			firstInstance := domain.UploadInstance{
				Idx: 0,
				Data: domain.User{
					Id:     1,
					Name:   "User 1",
					Status: constant.BLACKLIST,
				},
				Status: constant.FAILED,
				Reason: constant.STATUSNOTAVAILABLE,
			}

			secondInstance := domain.UploadInstance{
				Idx: 1,
				Data: domain.User{
					Id:     2,
					Name:   "User 2",
					Status: constant.BLACKLIST,
				},
				Status: constant.FAILED,
				Reason: errDummy.Error(),
			}

			// Assert
			if !reflect.DeepEqual(uploadInstances[0], firstInstance) {
				t.Errorf("False first upload instance when failed parsing occurs! got = %v, want %v", uploadInstances[0], firstInstance)
			}
			if !reflect.DeepEqual(uploadInstances[1], secondInstance) {
				t.Errorf("False second upload instance when failed parsing occurs! got = %v, want %v", uploadInstances[1], secondInstance)
			}
		}).Return(domain.NullResult(), nil)

		mockResultsRepo.On("CreateResult", mock.Anything).Return(int64(1), nil)

		u := NewUsecase(mockUsersRepo, mockResultsRepo, mockCsvBuilder)

		// Act
		_, err := u.UploadWhitelists(dummyCtx, file)

		// Assert
		assert.Nil(t, err)
	})

	t.Run("should not halt creating reports failed saving reports file", func(t *testing.T) {

		// Arrange
		contents := [][]string{
			{"1", "User 1", "2"},
			{"2", "User 2", "0"},
		}
		mockCsv := &mockcsv.ICsv{}
		mockCsv.On("GetContents").
			Return(contents, nil)

		mockCsvBuilder := &mockcsv.ICsvBuilder{}
		mockCsvBuilder.On("FromFile", mock.Anything).
			Return(mockCsv, nil)

		mockUsersRepo := &mockusersrp.IUsersRepo{}
		mockUsersRepo.On("GetUser", int64(1)).Return(domain.User{Id: 1, Name: "User 1", Status: constant.BLACKLIST}, nil)
		mockUsersRepo.On("GetUser", int64(2)).Return(domain.User{Id: 2, Name: "User 2", Status: constant.WHITELIST}, nil)
		mockUsersRepo.On("UpdateUsers", mock.Anything).Return([]error{errDummy, errDummy})

		mockResultsRepo := &mockresultsrp.IResultsRepo{}
		mockResultsRepo.On("SaveResult", mock.Anything).Run(func(args mock.Arguments) {
			// Stub
			arg0 := args.Get(0).(domain.Result)
			uploadInstances := arg0.Instances

			firstInstance := domain.UploadInstance{
				Idx: 0,
				Data: domain.User{
					Id:     1,
					Name:   "User 1",
					Status: constant.BLACKLIST,
				},
				Status: constant.FAILED,
				Reason: constant.STATUSNOTAVAILABLE,
			}

			secondInstance := domain.UploadInstance{
				Idx: 1,
				Data: domain.User{
					Id:     2,
					Name:   "User 2",
					Status: constant.BLACKLIST,
				},
				Status: constant.FAILED,
				Reason: errDummy.Error(),
			}

			// Assert
			if !reflect.DeepEqual(uploadInstances[0], firstInstance) {
				t.Errorf("False first upload instance when failed parsing occurs! got = %v, want %v", uploadInstances[0], firstInstance)
			}
			if !reflect.DeepEqual(uploadInstances[1], secondInstance) {
				t.Errorf("False second upload instance when failed parsing occurs! got = %v, want %v", uploadInstances[1], secondInstance)
			}
		}).Return(domain.NullResult(), errDummy)

		mockResultsRepo.On("CreateResult", mock.Anything).Return(int64(1), nil)

		u := NewUsecase(mockUsersRepo, mockResultsRepo, mockCsvBuilder)

		// Act
		_, err := u.UploadWhitelists(dummyCtx, file)

		// Assert
		assert.Nil(t, err)
	})

	t.Run("should be able to handle failed saving reports info to datastore", func(t *testing.T) { // Arrange
		contents := [][]string{
			{"1", "User 1", "2"},
			{"2", "User 2", "0"},
		}
		mockCsv := &mockcsv.ICsv{}
		mockCsv.On("GetContents").
			Return(contents, nil)

		mockCsvBuilder := &mockcsv.ICsvBuilder{}
		mockCsvBuilder.On("FromFile", mock.Anything).
			Return(mockCsv, nil)

		mockUsersRepo := &mockusersrp.IUsersRepo{}
		mockUsersRepo.On("GetUser", int64(1)).Return(domain.User{Id: 1, Name: "User 1", Status: constant.BLACKLIST}, nil)
		mockUsersRepo.On("GetUser", int64(2)).Return(domain.User{Id: 2, Name: "User 2", Status: constant.WHITELIST}, nil)
		mockUsersRepo.On("UpdateUsers", mock.Anything).Return([]error{errDummy, errDummy})

		mockResultsRepo := &mockresultsrp.IResultsRepo{}
		mockResultsRepo.On("SaveResult", mock.Anything).Run(func(args mock.Arguments) {
			// Stub
			arg0 := args.Get(0).(domain.Result)
			uploadInstances := arg0.Instances

			firstInstance := domain.UploadInstance{
				Idx: 0,
				Data: domain.User{
					Id:     1,
					Name:   "User 1",
					Status: constant.BLACKLIST,
				},
				Status: constant.FAILED,
				Reason: constant.STATUSNOTAVAILABLE,
			}

			secondInstance := domain.UploadInstance{
				Idx: 1,
				Data: domain.User{
					Id:     2,
					Name:   "User 2",
					Status: constant.BLACKLIST,
				},
				Status: constant.FAILED,
				Reason: errDummy.Error(),
			}

			// Assert
			if !reflect.DeepEqual(uploadInstances[0], firstInstance) {
				t.Errorf("False first upload instance when failed parsing occurs! got = %v, want %v", uploadInstances[0], firstInstance)
			}
			if !reflect.DeepEqual(uploadInstances[1], secondInstance) {
				t.Errorf("False second upload instance when failed parsing occurs! got = %v, want %v", uploadInstances[1], secondInstance)
			}
		}).Return(domain.NullResult(), errDummy)

		mockResultsRepo.On("CreateResult", mock.Anything).Return(int64(1), errDummy)

		u := NewUsecase(mockUsersRepo, mockResultsRepo, mockCsvBuilder)

		// Act
		got, err := u.UploadWhitelists(dummyCtx, file)

		// Assert
		assert.NotNil(t, err)
		assert.ObjectsAreEqualValues(EMPTYRESULTDTO, got)
	})

}
