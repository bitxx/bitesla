# Api Service

This is the Api service

Generated with

```
micro new github.com/bitxx/bitesla/service/service-api --namespace=bitesla --alias=api --type=api
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: bitesla.service.api
- Type: api
- Alias: api

## Dependencies

Micro services depend on service discovery. The default is consul.

```
# install consul
brew install consul

# run consul
consul agent -dev
```

## Usage

A Makefile is included for convenience

Build the binary

```
make build
```

Run the service
```
./api-api
```

Build a docker image
```
make docker
```