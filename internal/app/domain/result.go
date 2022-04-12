package domain

type UploadInstance struct {
	user   User
	status int64
}

type Result struct {
	Id        int64
	Instances []UploadInstance
}

func NullResult() Result {
	return Result{
		Id:        -1,
		Instances: make([]UploadInstance, 0),
	}
}
