# Bulk Upload System POC
<p align="center">
  <img src="https://raw.githubusercontent.com/egonelbre/gophers/master/vector/superhero/lifting-1TB.svg" />
</p>

## Why?
A lot of application need uploading system. When you only need to upload 200 - 300 data, using normal uploading system will not be a problem.

Problem arise when you are trying to upload a lot of data. There are various problem that can happen, such as context deadline or partial error on some row. There are various way to handle this problem, and the solution that we are going to build is only one way to solve it.

## What?
A simple bulk upload system to handle alot of data. It simply representing a fake blacklist uploading system, where operator can upload a bunch of user to be blacklisted. After the operator uploaded it, the operator then can see a report for each row (data) to check either the upload is succesful or not...

## Prerequisites
- Golang minimum v1.12 (https://golang.org/doc/install)
- Go Modules (https://blog.golang.org/using-go-modules)

## How to Run 
1. Clone the repository
```bash
git clone git@github.com:William9923/bulk-upload-poc.git
```
2. Setup the hooks
```bash
sh ./setup.sh
```

