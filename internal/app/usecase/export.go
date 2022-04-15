package usecase

import (
	"bytes"
	"context"
	"fmt"

	"github.com/William9923/bulk-upload-poc/pkg/csv"
)

var EMPTYFILE = *new(bytes.Buffer)

func (u Usecase) ExportUsers(ctx context.Context) (bytes.Buffer, error) {

	// 1. Fetch all users data
	users, err := u.usersRepo.GetUsers()
	if err != nil {
		return EMPTYFILE, err
	}

	// 2. Build csv data...
	header := []string{
		"Id",
		"Name",
		"Status",
	}

	contents := [][]string{}
	for _, user := range users {
		content := make([]string, len(header))
		content[0] = fmt.Sprintf("%d", user.Id)
		content[1] = user.Name
		content[2] = fmt.Sprintf("%d", user.Status)

		contents = append(contents, content)
	}

	csvFile := csv.New(header, contents)

	// 3. Export csv
	dataBuffer, err := csvFile.Export()
	if err != nil {
		return EMPTYFILE, err
	}

	return dataBuffer, nil
}
func (u Usecase) ExportResult(ctx context.Context, id int64) (bytes.Buffer, error) {

	// 1. Fetch all results data
	result, err := u.resultsRepo.GetResult(id)
	if err != nil {
		return EMPTYFILE, err
	}

	// 2. Build csv data...
	// save to file
	header := []string{
		"Id",
		"Name",
		"Status",
		"Upload Result",
		"Reason",
	}

	contents := [][]string{}
	for _, instance := range result.Instances {
		content := make([]string, len(header))
		content[0] = fmt.Sprintf("%d", instance.Data.Id)
		content[1] = string(instance.Data.Name)
		content[2] = fmt.Sprintf("%d", instance.Data.Status)
		content[3] = fmt.Sprintf("%d", instance.Status)
		content[4] = string(instance.Reason)

		contents = append(contents, content)
	}

	csvFile := csv.New(header, contents)

	// 3. Export csv
	dataBuffer, err := csvFile.Export()
	if err != nil {
		return EMPTYFILE, err
	}

	return dataBuffer, nil
}
