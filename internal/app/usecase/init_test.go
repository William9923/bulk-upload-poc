package usecase

import (
	"context"
	"fmt"
	"reflect"
	"testing"

	resultsrepo "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/results"
	mockresultsrp "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/results/mocks"
	usersrepo "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/users"
	mockusersrp "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/users/mocks"
	"github.com/stretchr/testify/assert"
)

var mockUsersRepo = mockusersrp.New(nil)
var mockResultsRepo = mockresultsrp.New(nil)
var dummyUsecase = &Usecase{
	usersRepo:   mockUsersRepo,
	resultsRepo: mockResultsRepo,
}
var dummyCtx = context.Background()
var dummyError = fmt.Errorf("Unit test mock error")

func TestNewUsecase(t *testing.T) {
	type args struct {
		usersRepo   usersrepo.IUsersRepo
		resultsRepo resultsrepo.IResultsRepo
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
			},
			want: dummyUsecase,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := NewUsecase(tt.args.usersRepo, tt.args.resultsRepo); !reflect.DeepEqual(got, tt.want) {
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
				mockUsersRepo.On("GetUsers").Return(nil, dummyError)

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
			if assert.Equal(t, tt.want, got) {
				t.Errorf("Usecase.ShowUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_ShowResults(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		u       Usecase
		args    args
		want    ResultsDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.ShowResults(tt.args.ctx)
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
