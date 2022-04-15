package domain

import "github.com/William9923/bulk-upload-poc/internal/app/constant"

type UploadInstance struct {
	Idx    int64
	Data   User
	Status int64
	Reason string
}

type Result struct {
	Id        int64            `json:"id"`
	Instances []UploadInstance `json:"-"`
	URL       string           `json:"url"`
}

func NullResult() Result {
	return Result{
		Id:        -1,
		Instances: make([]UploadInstance, 0),
	}
}

func NullUploadInstance(id int64, reason string) UploadInstance {
	return UploadInstance{
		Idx:    id,
		Data:   NullUser(),
		Status: constant.NOTPROCESSED,
		Reason: reason,
	}
}
