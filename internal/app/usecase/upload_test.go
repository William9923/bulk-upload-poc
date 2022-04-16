package usecase

import (
	"context"
	"io"
	"reflect"
	"testing"
)

func TestUsecase_UploadWhitelists(t *testing.T) {
	type args struct {
		ctx  context.Context
		file io.Reader
	}
	tests := []struct {
		name    string
		u       *Usecase
		args    args
		want    ResultDTO
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.UploadWhitelists(tt.args.ctx, tt.args.file)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.UploadWhitelists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.UploadWhitelists() = %v, want %v", got, tt.want)
			}
		})
	}
}
