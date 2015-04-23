test:
	ginkgo specs/vendors

run:
	go run main.go

dependencies:
	godep save
