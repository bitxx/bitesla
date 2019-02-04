GOPATH:=$(shell go env GOPATH)

build:
	cd service/api-gateway && make build
	cd service/service-user && make build
	cd service/service-exchange && make build