all: build

.PHONY: dep
dep:
	dep ensure

.PHONY: test
test: dep
	go test common/*.go

.PHONY: build
build: dep test
	mkdir -p bin
	go build -o bin/in in/*.go 
	go build -o bin/out out/*.go 
	go build -o bin/check check/*.go 

.PHONY: image
image:
	docker build -t rosti/concourse-build-number-resource .
	
