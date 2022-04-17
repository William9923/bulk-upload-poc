package usecase

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	"github.com/William9923/bulk-upload-poc/internal/app/domain"
	resultsrepo "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/results"
	mockresultsrp "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/results/mocks"
	usersrepo "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/users"
	mockusersrp "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/users/mocks"
	"github.com/William9923/bulk-upload-poc/pkg/csv"
	mockcsv "github.com/William9923/bulk-upload-poc/pkg/csv/mocks"
)

var mockUsersRepo = mockusersrp.New(nil)
var mockResultsRepo = mockresultsrp.New(nil)
var mockCsv = mockcsv.New(nil)
var mockCsvBuilder = mockcsv.NewBuilder(nil)
var dummyUsecase = &Usecase{
	usersRepo:   mockUsersRepo,
	resultsRepo: mockResultsRepo,
	csvBuilder:  mockCsvBuilder,
}
var dummyCtx = context.Background()
var errDummy = fmt.Errorf("Unit test mock error")

func TestNewUsecase(t *testing.T) {
	type args struct {
		usersRepo   usersrepo.IUsersRepo
		resultsRepo resultsrepo.IResultsRepo
		csvBuilder  csv.ICsvBuilder
	}
	tests := []struct {
		name string
		args args
		want IUsecase
	}{
		{
			name: "Success",
			args: args{
				usersRepo:   mockUsersRepo,
				resultsRepo: mockResultsRepo,
				csvBuilder:  mockCsvBuilder,
			},
			want: dummyUsecase,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUsecase(tt.args.usersRepo, tt.args.resultsRepo, tt.args.csvBuilder); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("NewUsecase() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_ShowUsers(t *testing.T) {

	tests := []struct {
		name    string
		want    UsersDTO
		wantErr bool
		mocks   func() *Usecase
	}{
		{
			name:    "Failed to fetch user data",
			want:    UsersDTO{},
			wantErr: true,
			mocks: func() *Usecase {

				mockUsersRepo := &mockusersrp.IUsersRepo{}
				mockUsersRepo.On("GetUsers").Return(nil, errDummy)

				return &Usecase{
					usersRepo: mockUsersRepo,
				}
			},
		},
		{
			name: "Success",
			want: UsersDTO{
				Users: []domain.User{
					domain.NullUser(),
					domain.NullUser(),
				},
			},
			wantErr: false,
			mocks: func() *Usecase {
				mockUsersRepo := &mockusersrp.IUsersRepo{}
				mockUsersRepo.On("GetUsers").Return([]domain.User{
					domain.NullUser(),
					domain.NullUser(),
				}, nil)

				return &Usecase{
					usersRepo: mockUsersRepo,
				}
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			u := dummyUsecase
			if tt.mocks != nil {
				u = tt.mocks()
			}

			got, err := u.ShowUsers(dummyCtx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.ShowUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(tt.want, got) {
				t.Errorf("Usecase.ShowUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_ShowResults(t *testing.T) {

	tests := []struct {
		name    string
		want    ResultsDTO
		wantErr bool
		mocks   func() *Usecase
	}{
		{
			name:    "Failed to fetch upload results",
			want:    ResultsDTO{},
			wantErr: true,
			mocks: func() *Usecase {

				mockResultsRepo := &mockresultsrp.IResultsRepo{}
				mockResultsRepo.On("GetResults").Return(nil, errDummy)

				return &Usecase{
					resultsRepo: mockResultsRepo,
				}

			},
		},
		{
			name: "Success",
			want: ResultsDTO{
				Results: []domain.Result{
					domain.NullResult(),
					domain.NullResult(),
				},
			},
			wantErr: false,
			mocks: func() *Usecase {

				mockResultsRepo := &mockresultsrp.IResultsRepo{}
				mockResultsRepo.On("GetResults").Return([]domain.Result{
					domain.NullResult(),
					domain.NullResult(),
				}, nil)

				return &Usecase{
					resultsRepo: mockResultsRepo,
				}

			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {

			u := dummyUsecase
			if tt.mocks != nil {
				u = tt.mocks()
			}
			got, err := u.ShowResults(dummyCtx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.ShowResults() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.ShowResults() = %v, want %v", got, tt.want)
			}
		})
	}
}
