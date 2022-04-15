package usecase

import (
	"context"
	"io"
	"strconv"

	"github.com/William9923/bulk-upload-poc/internal/app/constant"
	"github.com/William9923/bulk-upload-poc/internal/app/domain"
	"github.com/William9923/bulk-upload-poc/pkg/csv"
	"github.com/sirupsen/logrus"
)

func (u *Usecase) UploadWhitelists(ctx context.Context, file io.Reader) (ResultDTO, error) {

	dto := ResultDTO{
		Id:  domain.NullResult().Id,
		URL: domain.NullResult().URL,
	}

	logger := logrus.WithContext(ctx)
	uploadInstances, err := u.validateWhitelists(ctx, file)
	if err != nil {
		return dto, err
	}

	affectedUsers := make([]domain.User, len(uploadInstances))
	for i, instance := range uploadInstances {
		affectedUsers[i] = instance.Data
	}

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

	// build upload result reports...
	results := domain.Result{Instances: uploadInstances}
	results, err = u.resultsRepo.SaveResult(results)
	if err != nil {
		logger.Error("failed to save the result:", err)
		// TODO: should be able to rollback, but we use in memory db so its hard...
	}
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

	csvFile, err := csv.FromFile(file)
	if err != nil {
		logger.Error("unable to read file: ", err)
		return make([]domain.UploadInstance, 0), err
	}

	instances := make([]domain.UploadInstance, len(csvFile.Content))

	for i, data := range csvFile.Content {

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
