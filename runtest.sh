#!/bin/sh

OUTPUT_FILE_FOR_UNIT_TEST=test-coverage.out
readonly OUTPUT_FILE_FOR_UNIT_TEST

go test ./... -coverprofile=$OUTPUT_FILE_FOR_UNIT_TEST
#the ./... in the line above means: go test all subdirectory (/...) of current directory (./)
#for more details, run this from command line: go help package

#Reporting code coverage
echo "\n"
echo "Total coverage: "
go tool cover -func $OUTPUT_FILE_FOR_UNIT_TEST | grep total | awk '{print $3}'

#Add instruction to manually view coverage output
echo "\n"
echo "To view the coverage data, run this command from your terminal:"
echo "go tool cover -html=${OUTPUT_FILE_FOR_UNIT_TEST}"