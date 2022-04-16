package usecase

import (
	"reflect"
	"testing"

	resultsrepo "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/results"
	mockresultsrp "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/results/mocks"
	usersrepo "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/users"
	mockusersrp "github.com/William9923/bulk-upload-poc/internal/app/interface/repository/users/mocks"
)

var mockUsersRepo = mockusersrp.New(nil)
var mockResultsRepo = mockresultsrp.New(nil)

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
			want: &Usecase{
				usersRepo:   mockUsersRepo,
				resultsRepo: mockResultsRepo,
			},
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
