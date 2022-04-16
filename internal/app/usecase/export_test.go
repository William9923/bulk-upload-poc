package usecase

import (
	"bytes"
	"context"
	"reflect"
	"testing"
)

func TestUsecase_ExportUsers(t *testing.T) {
	type args struct {
		ctx context.Context
	}
	tests := []struct {
		name    string
		u       Usecase
		args    args
		want    bytes.Buffer
		wantErr bool
	}{}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.ExportUsers(tt.args.ctx)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.ExportUsers() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.ExportUsers() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestUsecase_ExportResult(t *testing.T) {
	type args struct {
		ctx context.Context
		id  int64
	}
	tests := []struct {
		name    string
		u       Usecase
		args    args
		want    bytes.Buffer
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := tt.u.ExportResult(tt.args.ctx, tt.args.id)
			if (err != nil) != tt.wantErr {
				t.Errorf("Usecase.ExportResult() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Usecase.ExportResult() = %v, want %v", got, tt.want)
			}
		})
	}
}
