all:
	go fmt .
	go fmt ./y24/**
	go build

test:
	go test ./y24/**
