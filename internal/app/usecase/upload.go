package usecase

import (
	"context"
	"io"
	"strconv"

	"github.com/William9923/bulk-upload-poc/internal/app/constant"
	"github.com/William9923/bulk-upload-poc/internal/app/domain"
	"github.com/sirupsen/logrus"
)

func (u *Usecase) UploadWhitelists(ctx context.Context, file io.Reader) (ResultDTO, error) {

	dto := ResultDTO{
		Id:  domain.NullResult().Id,
		URL: domain.NullResult().URL,
	}

	logger := logrus.WithContext(ctx)

	// 1. Validate Whitelist data & collect all affected user data...
	uploadInstances, err := u.validateWhitelists(ctx, file)
	if err != nil {
		return dto, err
	}
	affectedUsers := make([]domain.User, len(uploadInstances))
	for i, instance := range uploadInstances {
		affectedUsers[i] = instance.Data
	}

	// 2. Update user data using user repository & update the upload instance status...
	// TODO: add goroutine capability --> or could use message queue if data too big ...
	errs := u.usersRepo.UpdateUsers(affectedUsers)
	for i, err := range errs {
		uploadInstances[i].Status = constant.SUCCESS
		if err != nil {
			uploadInstances[i].Status = constant.FAILED

			if uploadInstances[i].Reason == constant.VALID {
				uploadInstances[i].Reason = err.Error()
			}
		}
	}

	// 3. Build reports file (notes: if using message queue, need different approach to build the reports...)
	results := domain.Result{Instances: uploadInstances}
	results, err = u.resultsRepo.SaveResult(results)
	if err != nil {
		logger.Error("failed to save the result:", err)
	}

	// 4. Save reports data so it can be used in the future...
	_, err = u.resultsRepo.CreateResult(results)
	if err != nil {
		logger.Error("unable to save upload results report: ", err)
		return dto, err
	}

	dto.Id = results.Id
	dto.URL = results.URL
	return dto, nil
}

func (u Usecase) validateWhitelists(ctx context.Context, file io.Reader) ([]domain.UploadInstance, error) {

	logger := logrus.WithContext(ctx)

	csvFile, err := u.csvBuilder.FromFile(file)
	if err != nil {
		logger.Error("unable to read file: ", err)
		return make([]domain.UploadInstance, 0), err
	}

	contents := csvFile.GetContents()
	instances := make([]domain.UploadInstance, len(contents))

	for i, data := range contents {

		idx := int64(i)
		userID, err := strconv.Atoi(data[0])
		if err != nil {
			instances[i] = domain.NullUploadInstance(idx, constant.FAILEDPARSING)
			continue
		}
		user, err := u.usersRepo.GetUser(int64(userID))
		if err != nil {
			instances[i] = domain.NullUploadInstance(idx, constant.USERNOTFOUND)
			instances[i].Data.Id = int64(userID)
			instances[i].Data.Name = data[1]
			continue
		}
		newStatus, err := strconv.Atoi(data[2])
		if err != nil || !isStatusValid(int64(newStatus)) {
			instances[i] = domain.NullUploadInstance(idx, constant.STATUSNOTAVAILABLE)
			instances[i].Data.Id = int64(userID)
			instances[i].Data.Name = data[1]
		}
		user.Status = int64(newStatus)

		instances[i] = domain.UploadInstance{
			Idx:    idx,
			Data:   user,
			Status: constant.NOTPROCESSED,
			Reason: constant.VALID,
		}
	}

	return instances, nil
}

func isStatusValid(status int64) bool {
	return status == constant.BLACKLIST || status == constant.WHITELIST
}
