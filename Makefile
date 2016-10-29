test:
	ginkgo specs/vendors

run:
	go run main.go

dependencies:
	godep save

build:
	go build

build_linux:
	GOOS=linux GOARCH=amd64 go build