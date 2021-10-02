-include main.env
export

GOPATH=$(shell go env GOPATH)

api:
	go run ./cmd/api/main.go