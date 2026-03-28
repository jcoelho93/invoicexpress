.PHONY: build run

build:
	go build -o invoicexpress cmd/main.go

run: build
	./invoicexpress
