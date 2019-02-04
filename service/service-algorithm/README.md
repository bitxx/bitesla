# Algorithm Service

This is the Algorithm service

Generated with

```
micro new github.com/jason-wj/bitesla/service/service-algorithm --namespace=bitesla --alias=algorithm --type=srv
```

## Getting Started

- [Configuration](#configuration)
- [Dependencies](#dependencies)
- [Usage](#usage)

## Configuration

- FQDN: bitesla.srv.algorithm
- Type: srv
- Alias: algorithm

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
./algorithm-srv
```

Build a docker image
```
make docker
```