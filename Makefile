GOPATH:=$(shell go env GOPATH)

build:
	cd service/service-api && make build
	cd service/service-user && make build
	cd service/service-exchange && make build
	cd service/service-strategy && make build