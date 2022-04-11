willi-repo=github.com/William9923

.PHONY: dep services cli nsq cron cover

cover: coverage-clear coverage-test coverage-total

coverage-clear:
	rm -rf coverage.out coverage.log

coverage-total:
	go tool cover -func coverage.out | grep total | awk '{print "\nTotal coverage: " $$3}'

coverage-test:
	go test -race `go list ./... | grep -v "cmd/" | grep -v "scripts/" | grep -v "internal/constant/"` -coverprofile=coverage.out

coverage-file:
	make coverage-folder folder=$(shell dirname ${file}) > /dev/null && go tool cover -html=coverage.out -o cover.html && cat cover.html | grep ${file}

coverage-folder:
	go test -race `go list ./... | grep ${folder}`  -coverprofile=coverage.out

dep: 
	GOPRIVATE=${willi-repo} go mod tidy

http:
	GOPRIVATE=${willi-repo} go run ./cmd/http/

cli:
	GOPRIVATE=${willi-repo} go run ./cmd/cli/

cron:
	GOPRIVATE=${willi-repo} go run ./cmd/cron/

line:
	GOPRIVATE=${willi-repo} go run ./cmd/line/